package main

import (
	"github.com/ggorockee/webserver-with-go/controllers"
	"github.com/ggorockee/webserver-with-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func setupRoutes(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g.POST("/users", controllers.CreateUser)

	// memo
	g.POST("/api/v1/memos", controllers.CreateMemo)
}

func main() {
	models.ConnectDB()
	r := gin.Default()
	setupRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
