package main

import (
	"log"
	"s3storage/internal/s3"
	"s3storage/pkg/configs"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	env := configs.NewEnvironment()

	level := env.AppConfig.LogLevel
	_ = configs.NewLogger(level)
	// logger.Debug("%+v", "AppConfig:", env.AppConfig)
	// logger.Debug("%+v", "AwsConfig:", env.AwsConfig)

	s3 := s3.NewS3Client(s3.AwsConfig{
		Endpoint:   env.AwsConfig.Endpoint,
		AccessKey:  env.AwsConfig.AccessKey,
		SecretKey:  env.AwsConfig.SecretKey,
		Region:     env.AwsConfig.Region,
		BucketName: env.AwsConfig.BucketName,
	})
	listBuckets := s3.ListBucket()

	log.Println("first page results")
	for _, object := range listBuckets {
		log.Printf("key=%s size=%d\n", aws.ToString(object.Key), *object.Size)
	}
	log.Printf("Total Count objects: %d", len(listBuckets))

	file := *listBuckets[0].Key
	log.Printf("Load file %s from S3", file)
	s3.DownloadFile(file, env.AwsConfig.OutputPath)
}
