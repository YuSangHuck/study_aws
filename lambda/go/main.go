package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func LongRunningHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	marshaledContext, err := json.Marshal(ctx)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}
	fmt.Println("Marshaled ctx: ", string(marshaledContext))

	lc, _ := lambdacontext.FromContext(ctx)
	marshaledLc, err := json.Marshal(lc)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}
	fmt.Println("Marshaled lc: ", string(marshaledLc))

	return events.APIGatewayProxyResponse{Body: "반갑수다", StatusCode: 200}, nil
	// deadline, _ := ctx.Deadline()
	// deadline = deadline.Add(-100 * time.Millisecond)
	// timeoutChannel := time.After(time.Until(deadline))

	// for {

	// 	select {

	// 	case <-timeoutChannel:
	// 		return "Finished before timing out.", nil

	// 	default:
	// 		log.Print("hello!")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// }
}

func main() {
	lambda.Start(LongRunningHandler)
}
