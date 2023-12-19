package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	dynamodb  "github.com/adapapooja/repository"
)

// Item represents the structure of the DynamoDB item
type Item struct {
	id   string `json:"id"`
	emailid string `json:"emailid"`
	FullName  string    `json:"fullname"`
	gender string `json:"gender"`
}

// createHandler for creating an item
func createHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newItem Item
	err := json.Unmarshal([]byte(request.Body), &newItem)
	if err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(newItem.id),
			},
			"emailid": {
				S: aws.String(newItem.emailid),
			},
			"FullName": {
				S:aws.String(newItem.FullName),
			},
		},
		TableName: aws.String(dynamodb.tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Error creating item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating item"}, nil
	}

	responseBody, _ := json.Marshal(newItem)
	return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(responseBody)}, nil
}

// readHandler for reading an item
func readHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	itemID := request.PathParameters["id"]

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		log.Printf("Error reading item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error reading item"}, nil
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Item not found"}, nil
	}

	var item Item
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("Error unmarshalling item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error reading item"}, nil
	}

	responseBody, _ := json.Marshal(item)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// updateHandler for updating an item
func updateHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	itemID := request.PathParameters["id"]

	var updatedItem Item
	err := json.Unmarshal([]byte(request.Body), &updatedItem)
	if err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(fmt.Sprintf("%d", updatedItem.Age)),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		UpdateExpression: aws.String("SET Age = :a"),
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		log.Printf("Error updating item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error updating item"}, nil
	}

	responseBody, _ := json.Marshal(updatedItem)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// deleteHandler for deleting an item
func deleteHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	itemID := request.PathParameters["id"]

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		log.Printf("Error deleting item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting item"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: ""}, nil
}
