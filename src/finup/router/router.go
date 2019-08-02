package router

import (
	. "finup/controller"
)
import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine  {
	var router=gin.Default()
	router.Use(Cors())
	router.Use(HeadersAuth())
	//router.Use(HeadersAuth())  //全局加入权限验证


	router.GET("/messages",SelectAllMessage)
	router.POST("/posttest",Posttest)
	router.POST("/getToken",GetToken)
	return router
}

