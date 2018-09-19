// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"go-grpc-app/action"
)

func main() {

	router := gin.Default()

	//api路由组
	api_router := router.Group("/api")
	api_user_router := api_router.Group("/user")
	{
		api_user_router.GET("/list", action.GetUsers)
	}

	router.Run()
}
