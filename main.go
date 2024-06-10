package main

import (
	"fmt"

	testcontroller "github.com/disky-business/tiny-url-svc/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/test-api", testcontroller.TestController)
	fmt.Println(">>>> Hello World <<<<")
}