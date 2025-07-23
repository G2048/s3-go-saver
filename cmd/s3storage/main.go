package main

import (
	"fmt"
	"log"
	"s3storage/cmd"
	"s3storage/internal/s3"
	"s3storage/pkg/configs"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	cmdArgs := cmd.NewCmdArgs()
	env := configs.NewEnvironment()

	level := env.AppConfig.LogLevel
	_ = configs.NewLogger(level)
	// logger.Debug("%+v", "AppConig:", env.AppConfig)
	// logger.Debug("%+v", "AwsConfig:", env.AwsConfig)

	s3 := s3.NewS3Client(s3.AwsConfig{
		Endpoint:   env.AwsConfig.Endpoint,
		AccessKey:  env.AwsConfig.AccessKey,
		SecretKey:  env.AwsConfig.SecretKey,
		Region:     env.AwsConfig.Region,
		BucketName: env.AwsConfig.BucketName,
	})

	if cmdArgs.List {
		listBuckets := s3.ListBucket()
		log.Println("first page results")
		for _, object := range listBuckets {
			fmt.Printf("key=%s size=%d\n", aws.ToString(object.Key), *object.Size)
		}
		fmt.Printf("\nTotal Count objects: %d\n", len(listBuckets))
	}

	// file := *listBuckets[0].Key
	// log.Printf("Load file %s from S3", file)
	// for _, object := range listBuckets {
	// s3.DownloadFile(*object.Key, env.AwsConfig.OutputPath)
	// }
}
