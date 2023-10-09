package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var choices []string

func AddChoices(c *gin.Context) {
	restaurant := c.PostForm("restaurant")
	choices = append(choices, restaurant)
	c.JSON(http.StatusOK, gin.H{"message": "Added Successfully"})
}

func GetRandomChoices(c *gin.Context) {
	if len(choices) == 0 {
		c.JSON(http.StatusOK, gin.H{"choice": "No restaurants added yet"})
		return
	}

	randomChoice := choices[rand.Intn(len(choices))]

	if randomChoice != "" {
		c.JSON(http.StatusOK, gin.H{"choices": randomChoice})
	}

}
