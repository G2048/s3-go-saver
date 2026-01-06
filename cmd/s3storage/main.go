package main

import (
	"fmt"
	"log"
	"os"
	"s3storage/cmd"
	"s3storage/configs"
	"s3storage/internal/s3"
)

func exist(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf("File '%s' is not found!", path)
		os.Exit(1)
	}
}

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
		listBuckets := *s3.ListBucket()
		log.Println("first page results")
		for _, object := range listBuckets {
			fmt.Printf("key=%s size=%d\n", object.Key, object.Size)
		}
		fmt.Printf("\nTotal Count objects: %d\n", len(listBuckets))
	}
	if cmdArgs.Upload != "" {
		exist(cmdArgs.Upload)
		fmt.Printf("Upload file '%s' to S3\n", cmdArgs.Upload)
		if err := s3.UploadFile(cmdArgs.Upload); err != nil {
			panic(err)
		}
	}
	if cmdArgs.Download != "" {
		fmt.Printf("Download file '%s' from S3\n", cmdArgs.Download)
		s3.DownloadFile(cmdArgs.Download, env.AwsConfig.OutputPath)
	}
	if cmdArgs.DowloadAll {
		listBuckets := *s3.ListBucket()
		log.Println("first page results")
		for _, object := range listBuckets {
			log.Printf("Load file %s from S3", object.Key)
			s3.DownloadFile(object.Key, env.AwsConfig.OutputPath)
		}
	}
	if cmdArgs.UploadAll != "" {
		exist(cmdArgs.UploadAll)
		fmt.Printf("Upload all files from '%s' to S3...\n", cmdArgs.UploadAll)
		if err := s3.UploadFiles(cmdArgs.UploadAll); err != nil {
			panic(err)
		}
	}
	if cmdArgs.Delete != "" {
		fmt.Printf("Delete file '%s' from S3\n", cmdArgs.Delete)
		if err := s3.DeleteFile(cmdArgs.Delete); err != nil {
			panic(err)
		}
	}
}
