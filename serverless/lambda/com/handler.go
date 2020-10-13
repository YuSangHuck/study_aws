package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	/*
		//
		// implement with gin-gonic
		//
		"log"
		"net/http"
		"os"
		"github.com/apex/gateway"
		"github.com/gin-gonic/gin"
	*/)

/*
//
// implement with gin-gonic
//
func test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "test1 lambda API"})
}
func test2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "test2 lambda API"})
}
*/

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	/*
		//
		// implement with gin-gonic
		//
		addr := ":3000"
		mode := os.Getenv("GIN_MODE")
		g := gin.New()
		g.GET("/test1", test1)
		g.GET("/test2", test2)

		if mode == "release" {
			log.Fatal(gateway.ListenAndServe(addr, g))
		} else {
			log.Fatal(http.ListenAndServe(addr, g))
		}
	*/

	//
	// implement with aws-sdk-lambda
	//
	lambda.Start(HandleLambdaEvent)
}
