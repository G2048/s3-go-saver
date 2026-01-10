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
func printS3(key string, size int64, keysOnly bool) {
	var str string
	if keysOnly {
		str = fmt.Sprintf("%s\n", key)
	} else {
		str = fmt.Sprintf("key=%s size=%d\n", key, size)
	}
	fmt.Printf(str)
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
			printS3(object.Key, object.Size, cmdArgs.KeysOnly)
		}
		fmt.Printf("\nTotal Count objects: %d\n", len(listBuckets))

	case cmdArgs.Upload != "":
		exist(cmdArgs.Upload)
		fmt.Printf("Upload file '%s' to S3\n", cmdArgs.Upload)
		if err := s3.UploadFile(cmdArgs.Upload); err != nil {
			panic(err)
		}

	case cmdArgs.Download != nil:
		fmt.Printf("Download file '%s' from S3\n", cmdArgs.Download)
		s3.DownloadFiles(cmdArgs.Download, env.AwsConfig.OutputPath, cmdArgs.IgnoreFullPath)

	case cmdArgs.DowloadAll:
		listBuckets := *s3.ListBucket()
		log.Println("first page results")

		var wg sync.WaitGroup
		for _, object := range listBuckets {
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Printf("Load file %s from S3", object.Key)
				s3.DownloadFile(object.Key, env.AwsConfig.OutputPath, false)
			}()
		}
		wg.Wait()

	case cmdArgs.UploadAll != "":
		exist(cmdArgs.UploadAll)
		fmt.Printf("Upload all files from '%s' to S3...\n", cmdArgs.UploadAll)
		if err := s3.UploadFiles(cmdArgs.UploadAll); err != nil {
			panic(err)
		}

	case cmdArgs.Delete != nil:
		fmt.Printf("Delete file '%s' from S3\n", cmdArgs.Delete)
		if err := s3.DeleteFiles(cmdArgs.Delete); err != nil {
			panic(err)
		}

	case cmdArgs.FuzzySearch != "":
		fmt.Printf("Fuzzy search files '%s' inside S3\n", cmdArgs.FuzzySearch)
		findedFiles, err := s3.FuzzySearchFile(cmdArgs.FuzzySearch)
		if err != nil {
			panic(err)
		}

		fmt.Println("\nFinded files:")
		for _, object := range findedFiles {
			printS3(object.Key, object.Size, cmdArgs.KeysOnly)
		}

	default:
		fmt.Printf("Command not found. For help using: -help option\n")
	}

	if cmdArgs.Time {
		elapsed := time.Since(start)
		fmt.Printf("Time of execution: %s\n", elapsed)
	}
}
