## [1.0.0] - 2026-01-11

### 🚀 Features

- *(pkg/s3/buckets.go)* Add to ListBucketOutput struct LastModified  field
- *(pkg/s3/files.go)* DownloadFiles() is async
- *(cmd/cli)* -delete flag recive many values
- *(cmd/cli/main.go)* Add flag -ignore-full-path; using with -download flag
- *(cmd/cli/main.go)* Replace method on s3.DownloadFiles() for -download flag
- *(pkg/s3/files.go)* Add DownloadFiles() method
- *(pkg/s3/files.go)* -fuzzy search by ignorcase
- *(cmd/cli)* Add flag -keys-only
- *(cmd/cli)* Add flag -fuzzy-search
- *(pkg/s3/files.go)* S3Client add FuzzySearchFile method
- *(cmd/cli/main.go)* Add flag -time for check time of execution programm
- *(cmd/s3-go-saver/main.go)* Parallel download for key -download-all
- *(cmd/s3storage/main.go)* Add upload all dir by flag -upload-all
- *(internal/s3/files.go)* Add S3Client.UploadFiles()
- *(pkg/tui/mainModel.go)* Add DeleteItem() method; delete object from disk
- *(internal/adapters/s3ListItems.go)* Implement for S3ListItems{} adapter the DownloadItems() method
- *(pkg/tui/mainModel.go)* Add to Storage interface the DownloadItems() method; embeding and call DownloadItems() method by Key.Download pressed
- *(internal/s3/s3.go)* Disable message of chekcsum validation skipped by dowlanding of object from S3 bucket
- *(configs/logs.go)* Add DisableLogs() function
- *(internal/adapters/s3ListItems.go)* S3ListItems implement the Storage{} interface
- *(pkg/tui/keymap.go)* Add Download keymap
- *(cmd/tui/main.go)* Draw content list items from S3 bucket in single tab
- *(pkg/tui/mainModel.go)* Chagne drawing list items for every tab one
- *(internal/adapters/s3ListItems.go)* ListBucket() return TabsItems type; split objectKey from S3 to folder and content
- *(cmd/tui/main.go)* Change list of content inside tui to list of s3 bucket entities
- *(internal/adapters/s3ListItems.go)* Add S3ListItems{} structure with ListBucket() method which returns the []list.Item list of structures for tui
- *(pkg/tui/items.go)* Add test list Items to main windows; in ModelTabs{}
- *(pkg/tui/mainModel.go)* Add a help tip strip
- *(cmd/tui/main.go)* Add run the tui ModelTabs
- *(pkg/tui/mainModel.go)* Add base sturcture ModelTabs are impelemented base tea model with base mehtods: Update, View, Init
- *(pkg/tui/windows.go)* Add structure Windows for drawing base window elements
- *(pkg/tui/keymap.go)* Add the ListKeyMap structure with basic keys: NextTab, PrevTab, HelpMenu, Exit
- *(pkg/tui/colours.go)* Add sturcture Colours for human-frendly color managment format
- *(internal/s3/buckets.go)* Add four features: upload, download, delte file and download all files
- *(cmd)* Add Upload and Download flags to CmdArgs structure
- *(cmd)* Add cli flag "--list" for list s3 buckets
- *(internal/s3/buckets.go)* Add to the S3Client.DownloadFile() output dir for saving s3 objects
- *(cmd/s3storage/main.go)* Add simple listing and download first file from S3 bucket
- *(internal/s3/buckets.go)* Add *Buckets methods for S3Client structure
- *(internal/s3/s3.go)* Add S3Client structure for manipulation S3
- *(pkg/configs/configs.go)* Add AppConfig and AwsConfig from .env
- *(pkg/configs/logs.go)* Add config for logger

### 🐛 Bug Fixes

- *(internal/adapters/s3ListItems.go)* Add withoutDir=false to S3.DownloadFile()
- *(pkg/s3/files.go)* Add return []ListBucketOutput
- *(cmd/cli/cmd.go)* Move code to single package args
- *(internal/s3/)* Move s3 package from internal to public pkg
- *(cmd/tui/main.go)* Change imports to s3-go-saver
- *(internal/adapters/s3ListItems.go)* Change imports to s3-go-saver
- *(cmd/cli/main.go)* -download-all: move waitGroup.Add(1) from gorutine
- *(cmd/cli/main.go)* -download-all: async download all files from bucket

### 💼 Other

- *(pkg/s3/files.go)* Incapsulate code from flag -download-all to DownloadAllFiles() method
- *(pkg/s3/buckets.go)* Rename ListBucketOutput to BucketObjects
- *(cmd/cli/args/args.go)* [**breaking**] Flag -download is []string
- *(cmd/cli/args/args.go)* Replace key -fuzzy-search to -fuzzy
- *(cmd/cmd.go)* Move to package main
- *(cmd/cli/main.go)* -time: split declaration of the start var from var assignment
- *(cmd/cli/main.go)* Switch-case: add \n for default help message
- *(cmd/)* Rename cmd/s3-go-saver to cmd/cli
- *(go.mod)* Go mod edit -go=1.25 && go mod tidy
- *(cmd/s3-go-saver/main.go)* Improve readability of swithc-case operator
- *(cmd/s3-go-saver/main.go)* Replace if to switch-case
- *(livereload.sh)* Rename building file from pug to livereloader
- *(go.mod)* Rename project to s3-go-saver
- *(go.mod)* Go mod tidy
- *(go.mod)* Add dependencies clipboard, bubbless, fuzzy;
- *(run.sh)* Add livereload.sh and run.sh like the air program for reloading app with code changes
- *(go.mod)* Add bubbletea - tui framework
- *(go.mod)* Update dependencies
- *(cmd/s3storage/main.go)* S3.DownloadFile to env.AwsConfig.OutputPath
- *(go.mod)* Add go-env, godotenv and aws packages
- *(git)* Add .gitignore

### 🚜 Refactor

- *(internal/adapters/s3ListItems.go)* Add const RootTab for s3 object keys witout dir in key
- *(pkg/tui/mainModel.go)* Rename Storage.DownloadItems() to Storage.DownloadItem()
- *(cmd/tui/main.go)* Disable logs messages and add to the S3ListItems adapter the OutputPath download point from env var
- *(pkg/tui/mainModel.go)* Add to input the NewModelTabs() the Storage{} interface
- *(pkg/tui/items.go)* Rename NewItems() func to NewTestItems()
- *(pkg/tui/mainModel.go)* NewModelTabs() on input take parameter []list.Item
- *(pkg/tui/items.go)* Add expliciet field Top and Desc to the Item{} structure
- *(pkg/tui/windows.go)* WindowDrawing.Style() align content by left side
- *(pkg/tui/mainModel.go)* ModelTabs to use Windows structure from variable "windows"
- *(pkg/tui/mainModel.go)* NewModelTabs() now accepts slices for Tabs and TabContent; added TestModelTabs() function to create a test instance
- *(pkg/tui/mainModel.go)* Change border figure for right active tab
- *(internal/s3/files.go)* Move the S3Client file methods to single file
- *(internal/s3/buckets.go)* Change output of S3Client.ListBucket() to []ListBucketOutpu; incapsulation of processing raw s3 objects.
- *(configs)* Move configs package to up
- *(internal/s3/s3.go)* Remove from AwsConfig.OutputPath

### 📚 Documentation

- *(README.md)* Add -keys-only; add -ignore-full-path; replace -fuzzy-search to -fuzzy
- *(README.md)* Add -fuzzy-search flag
- *(README.md)* Add -time description
- *(README.md)* Fix path in ## Build for Windows
- *(README.md)* Add command for building on Windows
- *(README.md)* Add command line arguments
- *(README.md)* Add command for build project
- Add README.md

### 🎨 Styling

- *(cmd/cli/main.go)* -download-all: add gap befor var wg sync.WaitGroup

### ⚙️ Miscellaneous Tasks

- *(pkg/s3/files.go)* Fix typo in S3Client.FuzzySearchFile()
- *(cmd/cli/main.go)* -time: fix message "Time of execution"
