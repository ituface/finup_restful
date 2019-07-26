package main

import (
	_ "./docs"
	"finup/router"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


// @title ituface
// @version 1.0

func main() {
	router:=router.InitRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	router.Run(":8882")
}
