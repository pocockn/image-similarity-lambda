package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	rekognitionLib "github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/pkg/errors"
)

type (
	// ShoutOutput holds the output from Rekognition and is what is stored within Dynamo.
	ShoutOutput struct {
		FaceMatches []*rekognitionLib.CompareFacesMatch `json:"face_matches" dynamodbav:"face_matches"`
		ShoutID     string                              `json:"shout_id" dynamodbav:"shout_id"`
	}
)

// Marshal converts values stored in a FreeTextAnalysis, into a DynamoDB
// item to be created within the store.
func (s ShoutOutput) Marshal() (*dynamodb.PutItemInput, error) {
	attributeValue, err := dynamodbattribute.MarshalMap(s)

	if err != nil {
		errors.Wrap(
			err,
			"Unable to convert dynamo.TextAnalysis struct into DynamoDB attribute value",
		)
	}

	return &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: imageSimilarityTableName(),
	}, nil
}

func imageSimilarityTableName() *string {
	t := "image_similarity"
	return &t
}
