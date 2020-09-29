package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "test1 lambda API"})
}
func test2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "test2 lambda API"})
}

func lambdaSDKHandler(ctx context.Context) (string, error) {
	return "반갑수다!", nil
}

func main() {
	addr := ":3000"
	mode := os.Getenv("GIN_MODE")
	g := gin.New()
	g.GET("/test1", test1)
	g.GET("/test2", test2)
	g.GET("/lambda-sdk", lambdaSDKHandler)

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(addr, g))
	} else {
		log.Fatal(http.ListenAndServe(addr, g))
	}
}
