package handlers

import (
	"encoding/json"
	"log"

	"github.com/adapapooja/repository"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Item represents the structure of the DynamoDB item
type Item struct {
	id   string `json:"id"`
	emailid string `json:"emailid"`
	FullName  string    `json:"fullname"`
	gender string `json:"gender"`
}
// createHandler for creating an item
func CreateHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
			"gender":{
				S:aws.String(newItem.gender),
			},
		},
		TableName: aws.String(repository.TableName),
	}

	_, err = repository.Svc.PutItem(input)
	if err != nil {
		log.Printf("Error creating item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating item"}, nil
	}

	responseBody, _ := json.Marshal(newItem)
	return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(responseBody)}, nil
}

// readHandler for reading an item
func ReadHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


	itemID := request.PathParameters["id"]

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		TableName: aws.String(repository.TableName),
	}

	result, err := repository.Svc.GetItem(input)
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
func UpdateHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


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
				S: aws.String( updatedItem.emailid),
			},
		},
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		UpdateExpression: aws.String("SET email id = :a"),
	}

	_, err = repository.Svc.UpdateItem(input)
	if err != nil {
		log.Printf("Error updating item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error updating item"}, nil
	}

	responseBody, _ := json.Marshal(updatedItem)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}

// deleteHandler for deleting an item
func DeleteHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


	itemID := request.PathParameters["id"]

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(itemID),
			},
		},
		TableName: aws.String(repository.TableName),
	}

	_, err := repository.Svc.DeleteItem(input)
	if err != nil {
		log.Printf("Error deleting item: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting item"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: ""}, nil
}
