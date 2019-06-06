package rekognition

import (
	"context"

	dynamoDBWrapper "github.com/pocockn/awswrappers/dynamodb"
	rekognitionWrapper "github.com/pocockn/awswrappers/rekognition"
)

type (
	// Client holds the Rekognition & Dynamo Client for interfacting
	// with the APIS.
	Client struct {
		DynamoDB    *dynamoDBWrapper.Client
		Rekognition *rekognitionWrapper.Client
	}
)

// Handle handles the request within Lambda, performing key phrase analysis and
// saving the analysed result in DynamoDB.
func (c Client) Handle(ctx context.Context) error {
	return nil
}
