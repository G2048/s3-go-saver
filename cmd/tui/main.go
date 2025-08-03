package main

import (
	"fmt"
	"os"
	"s3storage/configs"
	"s3storage/internal/adapters"
	"s3storage/internal/s3"
	"s3storage/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	env := configs.NewEnvironment()
	s3 := s3.NewS3Client(s3.AwsConfig{
		Endpoint:   env.AwsConfig.Endpoint,
		AccessKey:  env.AwsConfig.AccessKey,
		SecretKey:  env.AwsConfig.SecretKey,
		Region:     env.AwsConfig.Region,
		BucketName: env.AwsConfig.BucketName,
	})
	adapter := adapters.S3ListItems{S3: s3}
	// tabsItems := adapter.ListBucket()

	m := tui.NewModelTabs(&adapter)
	// m := tui.TestModelTabs()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
