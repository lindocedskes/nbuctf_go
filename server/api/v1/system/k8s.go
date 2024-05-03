package system

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	systemReq "github.com/lindocedskes/model/system/request"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"

	"go.uber.org/zap"
)

type K8sApi struct{}

// OpenContainer 传入 镜像地址如：lin088/w4terctf-2023:web001
func (s *K8sApi) OpenContainer(c *gin.Context) {
	var openImage systemReq.OpenImage
	err := c.ShouldBindJSON(&openImage) //绑定镜像地址
	if err != nil {
		global.GVA_LOG.Error("请求参数镜像地址错误!", zap.Error(err))
		response.FailWithMessage("请求参数镜像地址错误", c)
		return
	}
	userID := utils.GetUserID(c) //每个用户生成不同的pod
	containerIp, outPort, err2 := k8sService.OpenContainer(openImage, userID)

	if err2 != nil {
		if outPort == -1 {
			response.FailWithMessage("用户最多只能同时开启3个容器", c)
			return
		}
		global.GVA_LOG.Error("开启失败!", zap.Error(err2))
		response.FailWithMessage("开启失败", c)
		return
	}

	var closeImage systemReq.CloseImage
	closeImage.ImageAddr = openImage.ImageAddr
	// 设置一个定时器，在 2 小时后自动销毁容器,go程序不能中断
	time.AfterFunc(2*time.Hour, func() {
		flag, err := k8sService.CloseContainer(closeImage, userID)
		if err != nil {
			if flag != 1 && flag != 2 {
				global.GVA_LOG.Error("自动销毁容器失败!", zap.Error(err))
			}
		}
	})

	var resContainer = systemRes.ResContainer{
		ContainerIp: containerIp,
		OutPort:     outPort,
	}
	response.OkWithDetailed(resContainer, "开启成功", c)
}

// CloseContainer 关闭容器
func (s *K8sApi) CloseContainer(c *gin.Context) {
	var closeImage systemReq.CloseImage
	err := c.ShouldBindJSON(&closeImage) //绑定镜像地址
	if err != nil {
		global.GVA_LOG.Error("请求参数镜像地址错误!", zap.Error(err))
		response.FailWithMessage("请求参数镜像地址错误", c)
		return
	}
	userID := utils.GetUserID(c) //每个用户生成不同的pod
	flag, err2 := k8sService.CloseContainer(closeImage, userID)

	if err2 != nil {
		if flag == 1 {
			response.FailWithMessage("容器不存在,已关闭", c)
			return
		} else if flag == 2 {
			response.FailWithMessage("容器正在关闭，请稍后尝试", c)
			return
		} else {
			global.GVA_LOG.Error("关闭失败!", zap.Error(err2))
			response.FailWithMessage("关闭失败", c)
			return
		}

	}
	response.OkWithMessage("关闭成功", c)

}

// CheckContainer 查询容器运行状态
func (s *K8sApi) CheckContainer(c *gin.Context) {
	var checkImage systemReq.CheckImage
	err := c.ShouldBindJSON(&checkImage) //绑定镜像地址
	if err != nil {
		global.GVA_LOG.Error("请求参数镜像地址错误!", zap.Error(err))
		response.FailWithMessage("请求参数镜像地址错误", c)
		return
	}
	userID := utils.GetUserID(c) //每个用户生成不同的pod
	nodeIP, flag, err2 := k8sService.CheckContainer(checkImage, userID)

	if flag == 1 {
		response.FailWithMessage("容器不存在", c)
		return
	} else if flag == 2 {
		response.FailWithMessage("容器正在关闭", c)
		return
	} else if flag == 3 {
		global.GVA_LOG.Error("查询异常!", zap.Error(err2))
		response.FailWithMessage("查询异常", c)
		return
	} else if flag == 4 {
		response.FailWithMessage("容器正在创建", c)
		return
	} else if flag == 5 {
		response.FailWithMessage("容器正在启动...加载service", c)
		return
	}

	var resContainer = systemRes.ResContainer{
		ContainerIp: nodeIP,
		OutPort:     flag,
	}
	response.OkWithDetailed(resContainer, "容器正在运行", c)
	return
}

// CloseAllContainer 关闭所有容器
func (s *K8sApi) CloseAllContainer(c *gin.Context) {
	flag, err := k8sService.CloseAllContainer()

	if flag == 0 {
		response.FailWithMessage("已全部关闭", c)
		return
	} else if flag == 1 {
		response.OkWithMessage("关闭成功", c)
		return
	} else {
		global.GVA_LOG.Error("关闭失败!", zap.Error(err))
		response.FailWithMessage("关闭失败", c)
		return
	}

}
