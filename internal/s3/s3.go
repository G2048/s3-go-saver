package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Region     string
	BucketName string
	OutputPath string
}

type S3Client struct {
	AwsConfig
	s3 *s3.Client
}

func NewS3Client(awsConfig AwsConfig) *S3Client {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsConfig.Region),
		// config.WithCredentialsProvider(
		// credentials.NewStaticCredentialsProvider(
		// awsConfig.AccessKey, awsConfig.SecretKey, "",
		// ),
		// ),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client
	return &S3Client{
		AwsConfig: awsConfig,
		s3:        s3.NewFromConfig(cfg),
	}
}
