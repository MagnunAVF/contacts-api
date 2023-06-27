package db

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var Client *dynamodb.Client

func CreateClient(ctx context.Context) error {
	var cfg aws.Config
	var err error
	if os.Getenv("ENV") == "dev" {
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(os.Getenv("AWS_REGION")),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("accessKey", "secretKey", "token")),
			config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           "http://" + os.Getenv("HOST_IP") + ":8000",
					SigningRegion: os.Getenv("AWS_REGION"),
				}, nil
			})),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}

	if err != nil {
		return err
	}

	Client = dynamodb.NewFromConfig(cfg)

	return nil
}
