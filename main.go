package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	dynamoDBWrapper "github.com/pocockn/awswrappers/dynamodb"
	rekognitionWrapper "github.com/pocockn/awswrappers/rekognition"
	"github.com/pocockn/image-similarity-lambda/rekognition"
)

var (
	// Version is the tagged binary version.
	Version string
)

func main() {
	dynamoDBClient, err := createDynamoDBClient()
	if err != nil {
		log.Fatal(err)
	}

	client := rekognition.Client{
		Rekognition: rekognitionWrapper.NewClient(nil),
		DynamoDB:    dynamoDBClient,
	}

	lambda.Start(client.Handle)
}

func createDynamoDBClient() (*dynamoDBWrapper.Client, error) {
	return dynamoDBWrapper.NewClient(
		&dynamoDBWrapper.ClientConfig{},
		false,
		nil,
		nil,
	)
}
