package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (client *S3Client) ListBucket() []types.Object {
	output, err := client.s3.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(client.BucketName),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't list objects in bucket. Here's why: %s", err))
	}
	return output.Contents
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

func (client *S3Client) UploadFile(fileName string) error {
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't open file %v to upload. Here's why: %v", fileName, err))
		return err
	}
	defer file.Close()

	_, err = client.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't upload file %v to S3. Here's why: %v", fileName, err))
	}
	return err
}

func createFile(fullPath string) (*os.File, error) {
	var err error = nil
	parts := strings.Split(fullPath, "/")
	dirs := parts[:len(parts)-1]
	fileName := parts[len(parts)-1]

	slog.Info(fmt.Sprintf("Creating Directories: %v", dirs))
	slog.Info("Creating Directories")
	for _, path := range dirs {
		slog.Debug("Creating path: ", "path:", path)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			slog.Error(fmt.Sprintf("Couldn't create directory. Here's why: %s", err))
			return nil, err
		}
	}
	slog.Info(fmt.Sprintf("Creating file: %s", fileName))
	file, err := os.Create(fullPath)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't create file. Here's why: %s", err))
		return nil, err
	}
	return file, err
}
func (client *S3Client) DownloadFile(fileName string) error {
	result, err := client.s3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't get object from S3. Here's why: %s", err))
		return err
	}
	defer result.Body.Close()

	file, err := createFile(fileName)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't create file. Here's why: %s", err))
		return err
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		slog.Error(fmt.Sprintf("Cann'ot read body of s3 object. Causes: %s", err))
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't write file to disk. Here's why: %s", err))
	}
	return err
}
