package router

import (
	. "finup/controller"
)
import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine  {
	var router=gin.Default()


	router.GET("/messages",SelectAllMessage)
	return router
}

