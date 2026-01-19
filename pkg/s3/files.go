package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (client *S3Client) UploadFile(fileName, key string) error {
	var err error
	if key == "" {
		key = fileName
	}
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't open file %v to upload. Here's why: %v", fileName, err))
		return err
	}
	defer file.Close()

	_, err = client.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(client.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't upload file %v to S3. Here's why: %v", fileName, err))
	}
	return err
}
func (client *S3Client) UploadFiles(dirName string) error {
	var err error
	err = filepath.WalkDir(dirName, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			slog.Error(fmt.Sprintf("Couldn't open dir %s to upload. Here's why: %v", path, err))
			return err
		}

		fullPath, err := filepath.Abs(path)
		if err != nil {
			slog.Error(fmt.Sprintf("Couldn't open dir %s by full path to upload. Here's why: %v", fullPath, err))
			return err
		}

		if entry.Type().IsRegular() {
			slog.Info(fmt.Sprintf("Load into S3 the %s file...", path))
			file, err := os.Open(fullPath)
			if err != nil {
				slog.Error(fmt.Sprintf("Couldn't open file %s to upload. Here's why: %v", fullPath, err))
				return err
			}
			defer file.Close()

			_, err = client.s3.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket: aws.String(client.BucketName),
				Key:    aws.String(path),
				Body:   file,
			})
			if err != nil {
				slog.Error(fmt.Sprintf("Couldn't upload file %v to S3. Here's why: %v", fullPath, err))
				return err
			}

		}
		return nil
	})

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
func (client *S3Client) DownloadFiles(files []string, outputDir string, withoutDir bool) error {
	var (
		err error
		// TODO: change singnature for errors returning... (?)
		errors []error
		wg     sync.WaitGroup
	)

	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = client.DownloadFile(file, outputDir, withoutDir)
			if err != nil {
				slog.Warn(fmt.Sprintf("Couldn't download file %s from S3. Here's why: %v", file, err))
				errors = append(errors, err)
			}
		}()
	}
	wg.Wait()
	return err
}
func (client *S3Client) DownloadAllFiles(outputDir string) error {
	var (
		err error
		// TODO: change singnature for errors returning... (?)
		errors []error
		wg     sync.WaitGroup
	)

	listBuckets := *client.ListBucket()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't list object from S3. Here's why: %s", err))
		return err
	}
	slog.Info("First page results")
	for _, object := range listBuckets {
		wg.Add(1)
		go func() {
			defer wg.Done()
			slog.Info(fmt.Sprintf("Load file %s from S3", object.Key))
			err = client.DownloadFile(object.Key, outputDir, false)
			if err != nil {
				slog.Error(fmt.Sprintf("Couldn't download object from S3. Here's why: %s", err))
				errors = append(errors, err)
			}
		}()
	}
	wg.Wait()
	return err
}
func (client *S3Client) DownloadFile(fileName string, outputDir string, withoutDir bool) error {
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
	if withoutDir {
		_splits := strings.Split(fileName, "/")
		fileName = _splits[len(_splits)-1]
	}

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

func (s *S3Client) DeleteFiles(files []string) error {
	var err error
	for _, fileName := range files {
		err = s.DeleteFile(fileName)
		if err != nil {
			slog.Warn(fmt.Sprintf("Couldn't delete file %s from S3. Here's why: %v", fileName, err))
		}
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

func (client *S3Client) FuzzySearchFile(fileName string) ([]BucketObjects, error) {
	if fileName == "" {
		return nil, errors.New("Empty string for fuzzy search!")
	}
	r, err := regexp.Compile("(?i)" + fileName)
	if err != nil {
		slog.Error("Error by compile regex for fuzzing search! FileName: %s ; Error: %s", fileName, err)
		return nil, err
	}

	list, err := client.listBucket()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't list objects in bucket. Here's why: %s", err))
		return nil, err
	}

	var key string
	output := []BucketObjects{}
	for _, object := range list {
		key = aws.ToString(object.Key)
		if r.MatchString(key) {
			output = append(output, BucketObjects{aws.ToString(object.Key), *object.Size, *object.LastModified})
		}
	}
	return output, nil
}

var searchInFilesFormats = [...]string{"txt", "md"}

func hasExpectedFileFormat(file string) bool {
	for _, format := range searchInFilesFormats {
		if strings.HasSuffix(file, format) {
			return true
		}
	}
	return false
}

func (client *S3Client) InPlaceSearchFile(phrase string) ([]BucketObjects, error) {
	if phrase == "" {
		return nil, errors.New("Empty phrase for search in s3 files!")
	}

	r, err := regexp.Compile("(?i)" + phrase)
	if err != nil {
		slog.Error("Error by compile regex for inplace search! Phrase: %s ; Error: %s", phrase, err)
		return nil, err
	}

	list, err := client.listBucket()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't list objects in bucket. Here's why: %s", err))
		return nil, err
	}

	output := []BucketObjects{}
	var wg sync.WaitGroup
	for _, typeObject := range list {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Checking what searching in .txt and .md files
			if !hasExpectedFileFormat(aws.ToString(typeObject.Key)) {
				return
			}

			// Get S3 object
			slog.Info(fmt.Sprintf("Get object from S3: %s", aws.ToString(typeObject.Key)))
			object, err := client.s3.GetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: aws.String(client.BucketName),
				Key:    aws.String(aws.ToString(typeObject.Key)),
			})
			if err != nil {
				slog.Error(fmt.Sprintf("Couldn't get object from S3. Here's why: %s", err))
				// return nil, err
			}
			defer object.Body.Close()

			// Read object body
			fullBody := make([]byte, 0)
			buff := make([]byte, *typeObject.Size)
			for {
				n, err := object.Body.Read(buff)
				if n == 0 {
					break
				}
				if err != nil {
					slog.Error(fmt.Sprintf("Cannot read S3 object. Here's why: %s", err))
					// return nil, err
				}
				fullBody = append(fullBody, buff[:n]...)
			}
			// Concat file to single string
			var objectFullText string = string(fullBody)
			slog.Debug(objectFullText)

			// Final search text in file
			if r.MatchString(objectFullText) {
				output = append(output, BucketObjects{aws.ToString(typeObject.Key), *typeObject.Size, *object.LastModified})
			}
		}()
	}
	wg.Wait()
	return output, nil
}
