package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/fittest-individual", fittestIndividual)
	router.GET("/create-ga", createGA)

	err := router.Run("localhost:8080")
	if err != nil {
		panic("Failed to run router")
	}
}

func fittestIndividual(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func createGA(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
