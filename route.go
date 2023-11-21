package main

import (
	c "Mugen-Spiegel/findmo/app/controller/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/user/:id", c.UserController{}.Index)
		v1.POST("/user/", c.UserController{}.Create)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
