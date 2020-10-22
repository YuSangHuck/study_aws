package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	// marshaledRequest, err := json.Marshal(request)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	// }
	// fmt.Println("Marshaled request: ", string(marshaledRequest))

	// marshaledContext, err := json.Marshal(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	// }
	// fmt.Println("Marshaled ctx: ", string(marshaledContext))

	// lc, _ := lambdacontext.FromContext(ctx)
	// marshaledLc, err := json.Marshal(lc)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 500}, err
	// }
	// fmt.Println("Marshaled lc: ", string(marshaledLc))

	// if caller was 'serverless-plugin-warmup', return function
	if lc.ClientContext.Custom["source"] == "serverless-plugin-warmup" {
		fmt.Println("from serverless-plugin-warmup")
		return events.APIGatewayProxyResponse{Body: "from serverless-plugin-warmup", StatusCode: 202}, nil
	}

	// fmt.Println("from endpoint")

	start := time.Now()

	for i := 0; i < 1000000000; i++ {
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(elapsed.Nanoseconds())

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
