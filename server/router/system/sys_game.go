package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
)

type GameRouter struct{}

func (s *GameRouter) InitGameRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	gameRouter := Router.Group("game")
	gameApi := v1.ApiGroupApp.SystemApiGroup.GameApi
	{
		gameRouter.GET("list", gameApi.GameList)              // 获取题目列表
		gameRouter.POST("submitflag", gameApi.Submitflag)     // 提交flag
		gameRouter.GET("rank", gameApi.RankList)              // 排行榜
		gameRouter.POST("filedownload", gameApi.Filedownload) // 题目附件下载 by question_id
	}
	return gameRouter
}
