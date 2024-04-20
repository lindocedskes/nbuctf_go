package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
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
		gameType[v.QueType] = true //遍历存储所有存在题目类型
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
