# Necessary env variables
```
AWS_ENDPOINT_URL=https://storage.yandexcloud.net/
AWS_ACCESS_KEY=KEY
AWS_SECRET_KEY=SECRET_KEY
AWS_REGION=ru-central1
AWS_BUCKET_NAME=NAME
OUTPUT_DIR="output"
```

# Build Project
```
go build -v -o build/s3-cli cmd/s3-go-saver/main.go
```

## Build for Windows
```
GOOS=windows go build -v -o build/s3-cli.exe cmd/cli/main.go
```

# Command line arguments
```
Usage of ./build/s3-go-saver-cmd:
Program to list and download files from S3

  -delete string
    	Delete file from S3
  -download string
    	Download file from S3
  -download-all
    	Download all files from S3
  -list
    	List all files in bucket
  -upload string
    	Upload file to S3
  -upload-all string
    	Upload all files from specify directory to S3
```
