package router

import (
	. "finup/controller"
	_ "net/http"
)

import "github.com/gin-gonic/gin"

// Name will print hello name
// @Summary Print
// @Accept json
// @Tags Name
// @Security Bearer
// @Produce  json
// @Param name path string true "name"
// @Resource Name
// @Router /hello/{name} [get]
func InitRouter() *gin.Engine  {
	var router=gin.Default()


	router.GET("/messages",SelectAllMessage)
	return router
}

