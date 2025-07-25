package s3

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ListBucketOutput struct {
	Key  string
	Size int64
}

func (client *S3Client) listBucket() ([]types.Object, error) {
	output, err := client.s3.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(client.BucketName),
	})
	return output.Contents, err
}
func (client *S3Client) ListBucket() *[]ListBucketOutput {
	list, err := client.listBucket()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't list objects in bucket. Here's why: %s", err))
	}
	output := make([]ListBucketOutput, len(list))
	for _, object := range list {
		output = append(output, ListBucketOutput{aws.ToString(object.Key), *object.Size})
	}
	return &output
}

func (client *S3Client) CreateBucket() *s3.CreateBucketOutput {
	var exists *types.BucketAlreadyExists
	output, err := client.s3.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(client.BucketName),
	})
	if err != nil {
		if errors.As(err, &exists) {
			slog.Warn("Bucket already exists")
		} else {
			panic(err)
		}
	}
	return output
}

func (client *S3Client) DeleteBucket() *s3.DeleteBucketOutput {
	output, err := client.s3.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(client.BucketName),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't delete bucket. Here's why: %s", err))
	}
	return output
}
