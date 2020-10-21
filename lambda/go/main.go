package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func LongRunningHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
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
