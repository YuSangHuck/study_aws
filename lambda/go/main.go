package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func LongRunningHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// marshaledContext, err := json.Marshal(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	// }
	// fmt.Println("Marshaled ctx: ", string(marshaledContext))

	// lc, _ := lambdacontext.FromContext(ctx)
	// marshaledLc, err := json.Marshal(lc)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	// }
	// fmt.Println("Marshaled lc: ", string(marshaledLc))

	// fmt.Println("from endpoint")

	fmt.Println("from endpoint")
	start := time.Now()

	for i := 0; i < 1000000000; i++ {
	}
	timeElapsed := time.Now().Sub(start)
	// fmt.Println(timeElapsed.Nanoseconds())
	fmt.Println(timeElapsed)

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
