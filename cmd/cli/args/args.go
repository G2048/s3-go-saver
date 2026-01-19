package args

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type UploadArgs struct {
	Name string
	Path string
}
type CmdArgs struct {
	List           bool
	DowloadAll     bool
	Time           bool
	KeysOnly       bool
	IgnoreFullPath bool
	Version        bool
	UploadAll      string
	FuzzySearch    string
	InPlaceSearch  string
	Delete         stringSlice
	Download       stringSlice
	Upload         UploadArgs
}

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}
func (s *stringSlice) Set(value string) error {
	*s = append(*s, strings.Split(value, "\n")...)
	return nil
}

func NewCmdArgs() *CmdArgs {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Program to list and download files from S3\n\n")
		flag.PrintDefaults()
	}

	var download stringSlice
	flag.Var(&download, "download", "Download file from S3")
	var delete stringSlice
	flag.Var(&delete, "delete", "Delete file from S3")

	var list = flag.Bool("list", false, "List all files in bucket")
	var upload = flag.NewFlagSet("upload", flag.ExitOnError)
	var uploadPath = upload.String("path", "", "Upload file by name to S3")
	var uploadName = upload.String("name", "", "Custom name for file from S3")
	// var upload = flag.String("upload", "", "Upload file to S3")
	var uploadAll = flag.String("upload-all", "", "Upload all files from specify directory to S3")
	var downloadAll = flag.Bool("download-all", false, "Download all files from S3")
	var time = flag.Bool("time", false, "Add time of execution")
	var fuzzy = flag.String("fuzzy", "", "Fuzzy search files inside S3")
	var keys = flag.Bool("keys-only", false, "Print only keys without size")
	var ignoreFullPath = flag.Bool("ignore-full-path", false, "Ignore full path for downloading files. Using with only -download flag")
	var version = flag.Bool("version", false, "Print programm info")
	var inplace = flag.String("inplace", "", "Inplace fuzzy search in inside S3 files")

	flag.Parse()
	upload.Parse(os.Args[2:])
	uploadArgs := UploadArgs{*uploadName, *uploadPath}
	return &CmdArgs{
		Download:       download,
		Delete:         delete,
		Upload:         uploadArgs,
		List:           *list,
		DowloadAll:     *downloadAll,
		UploadAll:      *uploadAll,
		Time:           *time,
		FuzzySearch:    *fuzzy,
		KeysOnly:       *keys,
		IgnoreFullPath: *ignoreFullPath,
		Version:        *version,
		InPlaceSearch:  *inplace,
	}
}
