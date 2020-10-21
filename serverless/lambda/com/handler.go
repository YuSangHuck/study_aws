package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	//
	// implement with gin-gonic
	//
	// "log"
	// "net/http"
	// "os"
	// "github.com/apex/gateway"
	// "github.com/gin-gonic/gin"
)

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request: ", request)
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Println("Received body: ", request.Body)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	}
	fmt.Println("Marshaled request: ", string(marshaledRequest))
	marshaledContext, err := json.Marshal(ctx)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	}
	fmt.Println("Marshaled ctx: ", string(marshaledContext))

	log.Println("--- ctx & lc")
	log.Println(ctx)
	lc, _ := lambdacontext.FromContext(ctx)
	log.Println(lc.Identity.CognitoIdentityPoolID)
	log.Println(lc)
	marshaledLc, err := json.Marshal(lc)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	}
	fmt.Println("Marshaled lc: ", string(marshaledLc))

	return events.APIGatewayProxyResponse{Body: "반갑수다", StatusCode: 200}, nil
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
