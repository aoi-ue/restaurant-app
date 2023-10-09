package main

import (
	"net/http"
	"restaurant-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/add-choices", handlers.AddChoices)
	router.GET("/get-random-choices", handlers.GetRandomChoices)

	router.Run("127.0.0.1:8080")
}
