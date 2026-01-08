package main

import (
	"fmt"
	"log"
	"os"
	"s3-go-saver/cmd/cli/args"
	"s3-go-saver/configs"
	"s3-go-saver/pkg/s3"
	"sync"
	"time"
)

func exist(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf("File '%s' is not found!", path)
		os.Exit(1)
	}
}

func main() {

	cmdArgs := args.NewCmdArgs()
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

	var start time.Time
	if cmdArgs.Time {
		start = time.Now()
	}

	switch {
	case cmdArgs.List:
		listBuckets := *s3.ListBucket()
		log.Println("first page results")
		for _, object := range listBuckets {
			fmt.Printf("key=%s size=%d\n", object.Key, object.Size)
		}
		fmt.Printf("\nTotal Count objects: %d\n", len(listBuckets))

	case cmdArgs.Upload != "":
		exist(cmdArgs.Upload)
		fmt.Printf("Upload file '%s' to S3\n", cmdArgs.Upload)
		if err := s3.UploadFile(cmdArgs.Upload); err != nil {
			panic(err)
		}

	case cmdArgs.Download != "":
		fmt.Printf("Download file '%s' from S3\n", cmdArgs.Download)
		s3.DownloadFile(cmdArgs.Download, env.AwsConfig.OutputPath)

	case cmdArgs.DowloadAll:
		listBuckets := *s3.ListBucket()
		log.Println("first page results")

		var wg sync.WaitGroup
		for _, object := range listBuckets {
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Printf("Load file %s from S3", object.Key)
				s3.DownloadFile(object.Key, env.AwsConfig.OutputPath)
			}()
		}
		wg.Wait()

	case cmdArgs.UploadAll != "":
		exist(cmdArgs.UploadAll)
		fmt.Printf("Upload all files from '%s' to S3...\n", cmdArgs.UploadAll)
		if err := s3.UploadFiles(cmdArgs.UploadAll); err != nil {
			panic(err)
		}

	case cmdArgs.Delete != "":
		fmt.Printf("Delete file '%s' from S3\n", cmdArgs.Delete)
		if err := s3.DeleteFile(cmdArgs.Delete); err != nil {
			panic(err)
		}

	default:
		fmt.Printf("Command not found. For help using: -help option\n")
	}

	if cmdArgs.Time {
		elapsed := time.Since(start)
		fmt.Printf("Time of execution: %s\n", elapsed)
	}
}
