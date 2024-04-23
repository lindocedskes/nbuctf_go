package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/model/system/request"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
)

type GameApi struct{}

// 获取  未隐藏的题目列表 以及 当前用户的解题情况 和 所有题目类型 和 所有文件
func (s *GameApi) GameList(c *gin.Context) {
	userId := utils.GetUserID(c)               //从jwt获取用户ID
	GameList, err := gameService.GetGameList() //获取 未隐藏的题目列表
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	gameType := make(map[string]bool) //存储题目类型
	for index, v := range GameList {
		gameType[v.QueType] = true   //遍历存储所有存在题目类型
		GameList[index].QueFlag = "" //flag置空
		if gameService.IfSolved(v.ID, userId) {
			GameList[index].IfSolved = true //当前用户已解决该题
		} else {
			GameList[index].IfSolved = false //当前用户未解决该题
		}
	}

	allfiles, err := gameService.GetAllFiles() //获取所有文件
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(systemRes.GetGameListResponse{
		GameList: GameList,
		GameType: gameType,
		AllFiles: allfiles,
	}, "获取成功", c)
}

// 提交flag，记录，正确则加分
func (s *GameApi) Submitflag(c *gin.Context) {
	userId := utils.GetUserID(c) //从jwt获取用户ID
	var submit system.WrongGameRecord
	submit.UserId = userId
	_ = c.ShouldBindJSON(&submit) //绑定提交的flag, UserId , QuestionId
	if submit.SubmitFlag == "" {
		response.FailWithMessage("flag不能为空", c)
		return
	}

	if gameService.IfSolved(submit.QuestionId, userId) { //判断当前用户是否已解决该题
		response.FailWithMessage("已解决", c)
		return
	} else {
		judge, _ := gameService.JudgeFlag(submit.SubmitFlag, submit.QuestionId, userId) //提交flag，记录，正确则加分
		if judge {
			response.OkWithMessage("flag正确", c)
		} else {
			response.FailWithMessage("flag错误", c)
		}
	}

}

// 获取排名,降序 分数相同则按照提交时间升序
func (s *GameApi) RankList(c *gin.Context) {
	rankList, err := gameService.GetRankList() //获取排名
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rankList, "获取成功", c)
}

// 获取题目附件下载 by question_id
func (s *GameApi) Filedownload(c *gin.Context) {
	var submit system.WrongGameRecord
	err := c.ShouldBindJSON(&submit)
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return

	}
	questionId := submit.QuestionId

	files, err := gameService.GetFileById(questionId) //获取题目附件
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	//返回所有文件列表
	response.OkWithDetailed(files, "获取成功", c)
}

// 获取  所有题目列表 信息 和 所有文件
func (s *GameApi) GameListByAdmin(c *gin.Context) {

	GameList, err := gameService.GetGameListByAdmin() //获取 所有题目列表
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(GameList, "获取成功", c)
}

// 编辑题目，更新题目信息
func (s *GameApi) Editquestion(c *gin.Context) {
	var question system.Question
	err := c.ShouldBindJSON(&question) //绑定题目信息
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	err = gameService.EditQuestion(question) //编辑题目 byid
	if err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// 创建题目
func (s *GameApi) CreateQuestion(c *gin.Context) {
	var question system.Question
	err := c.ShouldBindJSON(&question) //绑定题目信息
	question.ID = 0                    //创建题目时，ID为空 uint 为无效值，gorm会自动创建
	question.IfSolved = false
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	question, err = gameService.CreateQuestion(question) //创建题目
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithDetailed(question, "创建成功", c)
}

// 删除题目 by question_id
func (s *GameApi) DeleteQuestion(c *gin.Context) {
	var question system.Question
	err := c.ShouldBindJSON(&question) //绑定题目信息
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	err = gameService.DeleteQuestion(question.ID) //删除题目 byid
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 添加题目附件 by question_id - file_id
func (s *GameApi) AddFile(c *gin.Context) {
	var questionFile request.QuestionFile
	err := c.ShouldBindJSON(&questionFile) //绑定题目附件信息
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	err = gameService.AddFile(questionFile.QuestionId, questionFile.FileId) //添加题目附件
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除题目附件 by question_id
func (s *GameApi) DeleteFiles(c *gin.Context) {
	var question system.Question
	err := c.ShouldBindJSON(&question) //绑定题目信息
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}
	err = gameService.DeleteFiles(question.ID) //删除题目附件 byid
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
