package main

import (
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", router))

}
