package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type GameRouter struct{}

func (s *GameRouter) InitGameRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	gameRouter := Router.Group("game")
	gameRouterWithRecord := Router.Group("gameadmin").Use(middleware.OperationRecord()) //记录关键操作
	gameApi := v1.ApiGroupApp.SystemApiGroup.GameApi
	{
		gameRouter.GET("list", gameApi.GameList)                     // 获取题目列表
		gameRouter.POST("submitflag", gameApi.Submitflag)            // 提交flag
		gameRouter.POST("rank", gameApi.RankList)                    // 排行榜，分页
		gameRouter.GET("submitScoreChart", gameApi.SubmitScoreChart) // 提交时间 - 得分 折线图表
		gameRouter.POST("filedownload", gameApi.Filedownload)        // 题目附件下载 by question_id

	}
	{
		gameRouterWithRecord.GET("list", gameApi.GameListByAdmin)           // 获取题目列表admin
		gameRouterWithRecord.POST("editquestion", gameApi.Editquestion)     // 编辑更新题目 by question_id
		gameRouterWithRecord.POST("createquestion", gameApi.CreateQuestion) // 创建题目
		gameRouterWithRecord.POST("deletequestion", gameApi.DeleteQuestion) // 删除题目 by question_id，并删除与之关联的file
		gameRouterWithRecord.POST("addfile", gameApi.AddFile)               // 添加题目附件 by question_id - file_id
		gameRouterWithRecord.POST("deletefiles", gameApi.DeleteFiles)       // 添加题目附件 by question_id
	}
	return gameRouter
}
