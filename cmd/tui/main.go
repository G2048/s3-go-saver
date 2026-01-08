package main

import (
	"fmt"
	"os"
	"s3-go-saver/configs"
	"s3-go-saver/internal/adapters"
	"s3-go-saver/pkg/s3"
	"s3-go-saver/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	env := configs.NewEnvironment()
	configs.DisableLogs()
	s3 := s3.NewS3Client(s3.AwsConfig{
		Endpoint:   env.AwsConfig.Endpoint,
		AccessKey:  env.AwsConfig.AccessKey,
		SecretKey:  env.AwsConfig.SecretKey,
		Region:     env.AwsConfig.Region,
		BucketName: env.AwsConfig.BucketName,
	})
	adapter := adapters.S3ListItems{S3: s3, DownloadDir: env.AwsConfig.OutputPath}
	m := tui.NewModelTabs(&adapter)
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
