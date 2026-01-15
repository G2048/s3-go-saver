package version

import (
	"fmt"
	"runtime/debug"
)

var (
	Application = "s3-go-saver"
	Version     = "dev"
	CommitSHA   = "N/A"
	BuildDate   = "N/A"
	GOARCH      = "N/A"
	GOOS        = "N/A"
	GoVersion   = "N/A"
	Sum         = "N/A"
	Compiler    = "N/A"
)

// HINT: go build -v -buildvcs
func GetVersionInfo() {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		GoVersion = buildInfo.GoVersion
		Version = buildInfo.Main.Version
		Sum = buildInfo.Main.Sum
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				CommitSHA = setting.Value
			case "vcs.time":
				BuildDate = setting.Value
			case "GOOS":
				GOOS = setting.Value
			case "GOARCH":
				GOARCH = setting.Value
			case "-compiler":
				Compiler = setting.Value
			}
		}
	}
}
func PrintVersionInfo() {
	fmt.Printf("Application: %s\n", Application)
	fmt.Printf("Go Version: %s\n", GoVersion)
	fmt.Printf("Go compiler: %s\n", Compiler)
	fmt.Printf("Application Version: %s\n", Version)
	fmt.Printf("Application Sum: %s\n", Sum)
	fmt.Printf("Platform: %s/%s\n", GOOS, GOARCH)
	fmt.Printf("Commit SHA: %s\n", CommitSHA)
	fmt.Printf("Build Date: %s\n", BuildDate)
}

func init() {
	GetVersionInfo()
}
