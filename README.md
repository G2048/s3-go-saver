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
make -B build
```
> Making binary file for linux and windows platforms

## Build for Windows
```
GOOS=windows go build -v -o build/s3-cli.exe cmd/cli/main.go
```

## With info about version
```
go build -v -buildvcs -o build/s3-cli ./cmd/cli/
```

# Command line arguments
```
Usage of s3-cli:
Program to list and download files from S3

  -delete value
    	Delete file from S3
  -download value
    	Download file from S3
  -download-all
    	Download all files from S3
  -fuzzy string
    	Fuzzy search files inside S3
  -ignore-full-path
    	Ignore full path for downloading files. Using with only -download flag
  -inplace string
    	Inplace fuzzy search in inside S3 files
  -keys-only
    	Print only keys without size
  -list
    	List all files in bucket
  -time
    	Add time of execution
  -upload string
    	Upload file to S3
  -upload-all string
    	Upload all files from specify directory to S3
  -version
    	Print programm info
```
