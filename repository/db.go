package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB client
var Svc *dynamodb.DynamoDB

// DynamoDB table name
const TableName = "crudoperations"

func init() {
	// Initialize the AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create DynamoDB client
	Svc = dynamodb.New(sess)
}
