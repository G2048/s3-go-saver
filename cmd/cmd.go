package cmd

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	List bool
}

func NewCmdArgs() *CmdArgs {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Program to list and download files from S3\n\n")
		flag.PrintDefaults()
	}
	var list = flag.Bool("list", false, "List all files in bucket")

	flag.Parse()
	return &CmdArgs{
		List: *list,
	}
}
