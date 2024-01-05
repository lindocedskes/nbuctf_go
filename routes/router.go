package routes

import (
	"ginVueBlog/utils"
	"net/http"
)
import "github.com/gin-gonic/gin"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1") //加前缀
	{
		router.GET("hello", func(c *gin.Context) { //get请求
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	//启动路由,utils.HttpPort 开放端口
	r.Run(utils.HttpPort)
}
