package configs

import (
	"fmt"
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	AppConfig struct {
		LogLevel string `env:"LOG_LEVEL,default=info"`
	}
	AwsConfig struct {
		Endpoint   string `env:"AWS_ENDPOINT_URL"`
		AccessKey  string `env:"AWS_ACCESS_KEY_ID"`
		SecretKey  string `env:"AWS_SECRET_KEY"`
		Region     string `env:"AWS_REGION"`
		BucketName string `env:"AWS_BUCKET_NAME"`
		OutputPath string `env:"OUTPUT_PATH,default=./output"`
	}
	Extras env.EnvSet
}

func NewEnvironment() Environment {
	var environment Environment
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}
	environment.Extras = es
	return environment
}
