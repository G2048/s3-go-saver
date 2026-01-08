package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	List       bool
	DowloadAll bool
	Time       bool
	UploadAll  string
	Download   string
	Upload     string
	Delete     string
}

func NewCmdArgs() *CmdArgs {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Program to list and download files from S3\n\n")
		flag.PrintDefaults()
	}
	var list = flag.Bool("list", false, "List all files in bucket")
	var upload = flag.String("upload", "", "Upload file to S3")
	var uploadAll = flag.String("upload-all", "", "Upload all files from specify directory to S3")
	var download = flag.String("download", "", "Download file from S3")
	var downloadAll = flag.Bool("download-all", false, "Download all files from S3")
	var delete = flag.String("delete", "", "Delete file from S3")
	var time = flag.Bool("time", false, "Add time of execution")

	flag.Parse()
	return &CmdArgs{
		List:       *list,
		Download:   *download,
		Upload:     *upload,
		DowloadAll: *downloadAll,
		Delete:     *delete,
		UploadAll:  *uploadAll,
		Time:       *time,
	}
}
