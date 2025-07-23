package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

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

	// Recursively create directories
	prevDir := ""
	slog.Debug(fmt.Sprintf("Creating Directories: %v", dirs))
	for _, path := range dirs {
		slog.Debug(fmt.Sprintf("Creating path: %s", path))
		prevDir = filepath.Join(prevDir, path)
		err = os.MkdirAll(prevDir, 0755)
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
func (client *S3Client) DownloadFile(fileName string, outputDir string) error {
	if outputDir == "" {
		outputDir = "."
	}
	result, err := client.s3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't get object from S3. Here's why: %s", err))
		return err
	}
	defer result.Body.Close()

	file, err := createFile(outputDir + "/" + fileName)
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
