package main

import (
	"net/http"
	"os"
	"restaurant-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	router.Static("/views", "./views")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/add-choices", handlers.AddChoices)
	router.GET("/get-random-choices", handlers.GetRandomChoices)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	router.Run(":" + port)
}
