package s3

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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

func (s *S3Client) DeleteFile(fileName string) error {
	_, err := s.s3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't delete object from S3. Here's why: %s", err))
		return err
	}
	return err
}
