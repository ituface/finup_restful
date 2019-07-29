package router

import (
	. "finup/controller"
)
import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine  {
	var router=gin.Default()
	router.Use(HeadersAuth())  //全局加入权限验证


	router.GET("/messages",SelectAllMessage)
	router.POST("/posttest",Posttest)
	return router
}

