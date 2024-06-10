package testcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context){
	c.IndentedJSON(http.StatusOK,
		gin.H{
			"message": "Hello World"})
}