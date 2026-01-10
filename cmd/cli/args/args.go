package args

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type CmdArgs struct {
	List        bool
	DowloadAll  bool
	Time        bool
	KeysOnly    bool
	UploadAll   string
	Upload      string
	Delete      string
	FuzzySearch string
	Download    stringSlice
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

	var list = flag.Bool("list", false, "List all files in bucket")
	var upload = flag.String("upload", "", "Upload file to S3")
	var uploadAll = flag.String("upload-all", "", "Upload all files from specify directory to S3")
	var downloadAll = flag.Bool("download-all", false, "Download all files from S3")
	var delete = flag.String("delete", "", "Delete file from S3")
	var time = flag.Bool("time", false, "Add time of execution")
	var fuzzy = flag.String("fuzzy", "", "Fuzzy search files inside S3")
	var keys = flag.Bool("keys-only", false, "Print only keys without size")

	flag.Parse()
	return &CmdArgs{
		Download:    download,
		List:        *list,
		Upload:      *upload,
		DowloadAll:  *downloadAll,
		Delete:      *delete,
		UploadAll:   *uploadAll,
		Time:        *time,
		FuzzySearch: *fuzzy,
		KeysOnly:    *keys,
	}
}
