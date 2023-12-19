package main

import (
	"context"

	"github.com/adapapooja/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}
// handler for handling CRUD operations
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		// return CreateHandler(request)
		return handlers.CreateHandler(request)
	case "GET":
		return handlers.ReadHandler(request)
	case "PUT":
		return handlers.UpdateHandler(request)
	case "DELETE":
		return handlers.DeleteHandler(request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid HTTP method"}, nil
	}
}