package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	rekognitionLib "github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/pkg/errors"
	dynamoDBWrapper "github.com/pocockn/awswrappers/dynamodb"
	rekognitionWrapper "github.com/pocockn/awswrappers/rekognition"
	"github.com/pocockn/image-similarity-lambda/rekognition/models"
)

const bucketName = "image-analysis-shouts"

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
func (c Client) Handle(ctx context.Context, shout models.ShoutInput) error {
	compareFacesInput := rekognitionLib.CompareFacesInput{
		SimilarityThreshold: aws.Float64(5.000000),
		SourceImage: &rekognitionLib.Image{
			S3Object: &rekognitionLib.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String("source_image.jpg"),
			},
		},
		TargetImage: &rekognitionLib.Image{
			S3Object: &rekognitionLib.S3Object{
				Bucket: aws.String(bucketName),
				Name:   aws.String("target_image.jpg"),
			},
		},
	}

	compareFaceOutput, err := c.Rekognition.CompareFaces(&compareFacesInput)
	if err != nil {
		return errors.Wrapf(err, "problem with comparing face with rekognition")
	}

	shoutOutput := models.ShoutOutput{
		ShoutID:     shout.ShoutID,
		FaceMatches: compareFaceOutput.FaceMatches,
	}

	_, err = c.DynamoDB.PutItem(shoutOutput)
	if err != nil {
		return errors.Wrapf(err, "problem putting item into dynamo")
	}

	return nil
}
