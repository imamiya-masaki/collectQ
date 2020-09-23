package main

import (
	"app/database"
	"app/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Origin",
			"Content-Type",
			"Content-Length",
		},
		AllowOrigins: []string{
			"*",
		},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})
	r.GET("/quetions", handler.ApiGetAllQuetions)
	r.Run()
}
