package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type answer struct {
	Question int    `json:"question"`
	Code     string `json:"code"`
}

func main() {
	testAPI()
	//router := gin.Default()
	//router.POST("/testAnswer", testAnswer)
	//router.GET("/create-ga", createGA)
	//
	//err := router.Run("localhost:8080")
	//if err != nil {
	//	panic("Failed to run router")
	//}
}

func testAPI() {
	const src = `func Bar(int1 int, int2 int) bool { if int1 + int2 == 5 { return true } else { return false } }`
	newAnswer := answer{Code: src, Question: 0}
	runAnswer(newAnswer)
}

func testAnswer(c *gin.Context) {
	var newAnswer answer
	// Call BindJSON to bind the received JSON to newAnswer
	if err := c.BindJSON(&newAnswer); err != nil {
		return
	}

	result := runAnswer(newAnswer)

	c.IndentedJSON(http.StatusOK, result)
}

func createGA(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
