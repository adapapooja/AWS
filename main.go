package main

import (
	"context"

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
		return createHandler(request)
	case "GET":
		return readHandler(request)
	case "PUT":
		return updateHandler(request)
	case "DELETE":
		return deleteHandler(request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid HTTP method"}, nil
	}
}