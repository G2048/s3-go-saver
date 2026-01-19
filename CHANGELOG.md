# Changelog

All notable changes to this project will be documented in this file.

## 1.2.1 - 2026-01-19

[e7dbd89](e7dbd8973fc1a5e72df3a7a59bb2e0fa92509528)...[002b4a1](002b4a13c4f7e17b397ad383aa404d5e322c91f5)

### Bug Fixes

- ListBucket(): right fill array of BucketObjects{} ([002b4a1](002b4a13c4f7e17b397ad383aa404d5e322c91f5))

### Documentation

- Update changelog for v1.2.0 ([6989b8f](6989b8f51654217f6d504c3e846769d9a8ae2f5b))
- Add the -upload key to -help print ([6ef27c7](6ef27c70d0de67e989527e5bb424fbe641487733))
- -help: add to -upload subflags -path and -name ([fc3c0f5](fc3c0f5d58ab2b4c4d041e052240f10799d08608))

## 1.2.0 - 2026-01-19

[97dc109](97dc1091411001305c46419604fc71d6e6c34199)...[e7dbd89](e7dbd8973fc1a5e72df3a7a59bb2e0fa92509528)

### Documentation

- Update changelog for v1.1.4 ([dc8b1a7](dc8b1a7fe960aec70084c68d40b0263a0be9de39))

### Features

- Change upload flag to under-flag; add the second subbcommands flag: -path and -name ([9e9d92e](9e9d92e2737de93dc6824c5b0e29c51256ade1a4))
- S3.Upload(): add choose a castom key for file uploading to S3 ([e7dbd89](e7dbd8973fc1a5e72df3a7a59bb2e0fa92509528))

## 1.1.4 - 2026-01-16

[d347e37](d347e371dd9c398e37e8fb207a42986ebc774af4)...[97dc109](97dc1091411001305c46419604fc71d6e6c34199)

### Bug Fixes

- Add info about commit, platform and version application ([97dc109](97dc1091411001305c46419604fc71d6e6c34199))

## 1.1.3 - 2026-01-16

[12bf4ea](12bf4ea87cc21b5c405ff04664a90e0d9c99881e)...[d347e37](d347e371dd9c398e37e8fb207a42986ebc774af4)

### Bug Fixes

- Build: replace MAINGO to MAIN_GO var ([d347e37](d347e371dd9c398e37e8fb207a42986ebc774af4))

## 1.1.2 - 2026-01-16

[8b2973c](8b2973c2b14c84f49ac4a0b3e27d3466c22f1357)...[12bf4ea](12bf4ea87cc21b5c405ff04664a90e0d9c99881e)

### Documentation

- Update changelog for v1.1.1 ([12bf4ea](12bf4ea87cc21b5c405ff04664a90e0d9c99881e))

## 1.1.1 - 2026-01-16

[d4e020c](d4e020c6b5560c7926eae60d6e95c040b75fc617)...[8b2973c](8b2973c2b14c84f49ac4a0b3e27d3466c22f1357)

### Documentation

- Update changelog for v1.1.0 ([93a78b9](93a78b950a810979f8fdb57126bcf95932695fbc))
- Update "# Build Project" section ([8b2973c](8b2973c2b14c84f49ac4a0b3e27d3466c22f1357))

### Miscellaneous Tasks

- Remove unnecessary comments ([4a573a6](4a573a6f3e18dca0a37736f43c115b30d49e41bb))

### Build

- Declare building steps ([ed23942](ed23942f463e4fd6c4507d0d0d566b77dfaf3a69))

### Ref

- Add type LogLevel; add consts for log levels; change input signature for NewLogger() func ([f8d6632](f8d6632e3c4d30857ac91b73512285358a595f5a))
- Add debug AwsConfig and AppConfig structures ([38dad7a](38dad7a49399a3c1bb5a784e665c58d4c9d8a617))

## 1.1.0 - 2026-01-15

[dcb2563](dcb2563de5b1567bc818ba8a3c2544872223df43)...[d4e020c](d4e020c6b5560c7926eae60d6e95c040b75fc617)

### Bug Fixes

- Logs field with application name and version info get dynammically ([89f4149](89f4149500945de078527e90a79c54b6b4179b00))

### Documentation

- Add "## With info about version" ([bc649c6](bc649c6657295022ab02ed1ec917dbd83dc9703c))
- Add CHANGELOG.md ([98d5bcd](98d5bcd54c432c78e032187aa13d20b996ca31ae))
- Add -verison and -inplace flags ([a34ca5d](a34ca5df0d8deead4aeb1035d40e67e4774c487e))

### Features

- Add print_version_info() function ([af9f4df](af9f4df07098ecdbcaf177552f661783506ae714))
- Add flag -version ([34ac5b9](34ac5b94c8f13a2e171c2f0e7da96ad9c167bafd))
- Print compiler info ([2c17a27](2c17a27e885777ab777c7f466bf867d8531d14d5))
- Add S3Client.InPlaceSearchFile() method for finding phrase inside s3 file ([122d14b](122d14b660fd7a238bf86c71c6372b9a6be1c2c7))
- Add flag -inplace for finding phrase inside s3 file ([7a1feac](7a1feac76acc45f4d3b3ef456e0ef16b26aea2ab))

### Ref

- Replace "Application building" to "Platform" ([f6d6fc1](f6d6fc1a1d0ebd23b3f41dc1346cb95e41ab8fbd))
- Move code in single package "version"; add const Application ([1121b67](1121b673aa6efb0a6494e24b35a8fc762301cc43))
- Change order of output; change print name of Verions and Sum info ([d4e5e2b](d4e5e2b40c5f37f93d707bfc8a49e6eb19b2f35e))

## 1.0.0 - 2026-01-11

### Bug Fixes

- -download-all: async download all files from bucket ([4299f82](4299f82b95e5cbd5f5337de672fa8731ba86975d))
- -download-all: move waitGroup.Add(1) from gorutine ([ed0f0ba](ed0f0ba6e0c029e40654f537a41c6a174d034dab))
- Change imports to s3-go-saver ([4f23c78](4f23c788e6cb104bbb73e4d09290a2447bd0ecf7))
- Change imports to s3-go-saver ([68a85c2](68a85c2c36b591efb6426b93528959e78bce9a70))
- Move s3 package from internal to public pkg ([2c420af](2c420af40e9fae5dc5b7baa21c4e812d6b6a2125))
- Move code to single package args ([44afc2d](44afc2de253503805631f8645960a5ba2f3ffb5e))
- Add return []ListBucketOutput ([8b76892](8b76892abdd758adc531ad5e45d109770ed271f1))
- Add withoutDir=false to S3.DownloadFile() ([dcb2563](dcb2563de5b1567bc818ba8a3c2544872223df43))

### Documentation

- Add README.md ([ecd9dbf](ecd9dbf2c448c9db075ab2315be784a1e6e6ccfd))
- Add command for build project ([c53d825](c53d825e947add63c6b64c890c5b4cec5bc56bf6))
- Add command line arguments ([25f72df](25f72df1dcd65c9c86956ee631716e5e624b5cd0))
- Add command for building on Windows ([9914510](991451015afa24f6d8608a62bbd9040455e9f933))
- Fix path in ## Build for Windows ([9348dbd](9348dbdd980ef15b887ecdba1ea4a5b625253ba3))
- Add -time description ([dd2a3d6](dd2a3d6a2906c864ad0c5019f5a1756506f3eaeb))
- Add -fuzzy-search flag ([b8ad50b](b8ad50bfd40a25034bb6e7cf021aded9388c24c2))
- Add -keys-only; add -ignore-full-path; replace -fuzzy-search to -fuzzy ([590696c](590696c46d6f88f4cdc652bdecd347b165d2164c))

### Features

- Add config for logger ([5fe1bff](5fe1bff1f75bfdcb7c7f346ed5023ba7635e270f))
- Add AppConfig and AwsConfig from .env ([0be4dfe](0be4dfeb604b74503053f741dda514c0f5e6a447))
- Add S3Client structure for manipulation S3 ([31c484d](31c484db74e2a1a952420e65a1635b4bf3034b9d))
- Add *Buckets methods for S3Client structure ([6b4f8da](6b4f8dac42ab65a259450f20c7e0be3d1ad2d614))
- Add simple listing and download first file from S3 bucket ([50723e3](50723e3730c8e0269a64de9183b46c81849713f0))
- Add to the S3Client.DownloadFile() output dir for saving s3 objects ([a6db13e](a6db13ecc5f62718eefcb9c2056efc058acfdaac))
- Add cli flag "--list" for list s3 buckets ([5de0e84](5de0e84384621f5fd03a6643cae083709a5c5f83))
- Add Upload and Download flags to CmdArgs structure ([9aee95a](9aee95aee0be7a7df03d99e57a6aa0f2ab03436b))
- Add four features: upload, download, delte file and download all files ([374a0ef](374a0efd5b0c508ec96a734e95186656ad71b47e))
- Add sturcture Colours for human-frendly color managment format ([ca1085e](ca1085e8ceaad39506f7ad82169c232250aff569))
- Add the ListKeyMap structure with basic keys: NextTab, PrevTab, HelpMenu, Exit ([4c392ba](4c392ba4e5e01928caef27714f24f61ddfe20518))
- Add structure Windows for drawing base window elements ([5ab6974](5ab6974893fd9fdef8bbc873ee96809e092e539b))
- Add base sturcture ModelTabs are impelemented base tea model with base mehtods: Update, View, Init ([88eaaca](88eaacae5b1c9ae394e6b983e9ae989856c2f6c9))
- Add run the tui ModelTabs ([7d1fc73](7d1fc730946b914e55fc769d03dba670d8caf43f))
- Add a help tip strip ([2d688af](2d688afcaf59a82e554c933ee5f4eec1d9cc5142))
- Add test list Items to main windows; in ModelTabs{} ([498c8e4](498c8e4032e1a4c2d185850d843688f668ffd35e))
- Add S3ListItems{} structure with ListBucket() method which returns the []list.Item list of structures for tui ([4fd487e](4fd487efe01e30930dd08b27ecee58fcd050b3a6))
- Change list of content inside tui to list of s3 bucket entities ([8c61c7c](8c61c7c4f0e29d5a0306315bf2ffd68defd7a69a))
- ListBucket() return TabsItems type; split objectKey from S3 to folder and content ([de9e460](de9e46022ee127694ab347677b6c2b506ba92b05))
- Chagne drawing list items for every tab one ([545fae1](545fae16d19a76c2160d9e18f045014c52dd128e))
- Draw content list items from S3 bucket in single tab ([a0162d5](a0162d54f49d91b1a8e5f0f96f97724ff61a0698))
- Add Download keymap ([c39d8a5](c39d8a586626043e14395cc6c36e639abf04c17e))
- S3ListItems implement the Storage{} interface ([bff679f](bff679f3379f32926e5f8f64f10b7d50771476d3))
- Add DisableLogs() function ([225cb66](225cb66d7c419bf55c8f5713e19037fced7fe7d5))
- Disable message of chekcsum validation skipped by dowlanding of object from S3 bucket ([0f2dbd9](0f2dbd91a0c5807511da35b3daa308ce9dd4b53d))
- Add to Storage interface the DownloadItems() method; embeding and call DownloadItems() method by Key.Download pressed ([d9c6027](d9c6027335e029687d9214a8672fcdbf66afe9c4))
- Implement for S3ListItems{} adapter the DownloadItems() method ([c02d5c4](c02d5c435f6f0fbfe2f2b3962e93c1be132a4ad2))
- Feat(pkg/tui/mainModel.go): add marker for downloaded files; change marker if file downloaded ([cab6e32](cab6e32668964393660c2adde74a6182157c5e6a))
- Add DeleteItem() method; delete object from disk ([bdd5e29](bdd5e29a847241930689ebecdf108d79c11317a8))
- Add S3Client.UploadFiles() ([39c4805](39c4805cca49609d9121d55ee152dd2da263526f))
- Add upload all dir by flag -upload-all ([b868757](b868757e02a4d8feccd48943903fc6d024a2e693))
- Parallel download for key -download-all ([17ded97](17ded978bdf4373f89bb53ebec9a9f356651035a))
- Add flag -time for check time of execution programm ([d4192c0](d4192c063a0d767cfa52e8ab028bb0adf1e50167))
- S3Client add FuzzySearchFile method ([a35fe3e](a35fe3e71f004fdbd32478b5d1e9cda9c6df5516))
- Add flag -fuzzy-search ([5d570f4](5d570f4ec23b1b8599dd1ad7f15ab34514c1eebf))
- Add flag -keys-only ([19454fe](19454fe569d3ab36f07594e5339bcfde095311e2))
- -fuzzy search by ignorcase ([612406b](612406b510e34656cf460a13f1ef3f09c43e2697))
- Add DownloadFiles() method ([0aa2bbd](0aa2bbd5fd2a57e05cd17875d24a26a176666e10))
- Replace method on s3.DownloadFiles() for -download flag ([cfbbd5f](cfbbd5f5d00ce493aac9f8ce956e0c82356912d6))
- Add flag -ignore-full-path; using with -download flag ([4855432](4855432c3a9857a11bf0e0ba2175dcba19439421))
- -delete flag recive many values ([a58df5d](a58df5d1a004f189669ddad476b2cc36e3541f9e))
- DownloadFiles() is async ([6bced3b](6bced3bbbccc1708a622e8f2fa64e271686e3ffa))
- Add to ListBucketOutput struct LastModified  field ([0caa5bd](0caa5bd177e724722a3f29bc925149f111da4e36))

### Miscellaneous Tasks

- -time: fix message "Time of execution" ([2b2447e](2b2447ef65c6a67cde44692a82665fde860c4d7f))
- Fix typo in S3Client.FuzzySearchFile() ([6c3fd87](6c3fd8731833f57196674bbae2b613c886ca7a11))

### Refactor

- Remove from AwsConfig.OutputPath ([42e5726](42e5726002661449f4d52a826ece4c88ffea4c8a))
- Move configs package to up ([7a99d10](7a99d108a39f8868ef4f30eeb901a33e121e9b18))
- Change output of S3Client.ListBucket() to []ListBucketOutpu; incapsulation of processing raw s3 objects. ([3e96c8a](3e96c8a02fe2b709de8e23943716b5d02a28c44d))
- Move the S3Client file methods to single file ([de8f1a2](de8f1a217fb9670010c507c6a1661b3d26cf3b40))
- Change border figure for right active tab ([a8cd348](a8cd34810b0b0002409d3e6a574108704b98e19a))
- NewModelTabs() now accepts slices for Tabs and TabContent; added TestModelTabs() function to create a test instance ([37e4868](37e48688be3995815dab6a6303e65461e8c6649c))
- ModelTabs to use Windows structure from variable "windows" ([a1dd758](a1dd758b0190cbe15f2be5d75ad17e6e6cd5e748))
- WindowDrawing.Style() align content by left side ([0fe8beb](0fe8beb3e907a3c3de0d733e6590a6782ba1a49c))
- Add expliciet field Top and Desc to the Item{} structure ([4bee465](4bee465aea7854f49433d80436754974bb8034a8))
- NewModelTabs() on input take parameter []list.Item ([455134e](455134e51d5523820c9cc680f62292a00185fcf7))
- Rename NewItems() func to NewTestItems() ([c4a33c1](c4a33c1eb5aebd81772d06850b26dee9db34c184))
- Add to input the NewModelTabs() the Storage{} interface ([0648e5b](0648e5b31c44179d4ab2f6775fbe25b4dfc6f2fd))
- Disable logs messages and add to the S3ListItems adapter the OutputPath download point from env var ([975e22c](975e22cfef62bad7bdd873f3d939b8063c3ac4a1))
- Rename Storage.DownloadItems() to Storage.DownloadItem() ([542fd77](542fd77c583d98cc33a7366f8775ec819deb487d))
- Add const RootTab for s3 object keys witout dir in key ([a1f259a](a1f259adf8574f646fa7707b8ff144962a8f3a37))

### Styling

- -download-all: add gap befor var wg sync.WaitGroup ([c0251b3](c0251b36fbc3c3435b28abd4c7f67306947e61a0))

### Build

- Add .gitignore ([3853560](38535603320d44a96fcb2aaca46619e577c470e6))
- Add go-env, godotenv and aws packages ([1fdfe79](1fdfe797805b4dc9c3c04b4b186dd7ec64a8cc38))
- Update dependencies ([dcf4a0f](dcf4a0f8e22b62610a9ec931786b1161099262c2))
- Add bubbletea - tui framework ([d12d1a9](d12d1a91e817f10671b512f8028a6d15f5468bd1))
- Add livereload.sh and run.sh like the air program for reloading app with code changes ([9e4a06f](9e4a06f201a0528058b2665182fbbf44b9448f2d))
- Add dependencies clipboard, bubbless, fuzzy; ([43da471](43da47162d7e1ec5169fc39f446114896d1572cb))
- Go mod tidy ([d716a83](d716a83dfda0b1fc0ccfe3443b81a9e010a58a4f))
- Go mod edit -go=1.25 && go mod tidy ([589589b](589589b6e97d27c52fed5e5d870bf33e343cda3b))

### Ref

- S3.DownloadFile to env.AwsConfig.OutputPath ([863d2ee](863d2eef5c39fe5e311d8a44d96215585c9f6726))
- Rename project to s3-go-saver ([2aa689b](2aa689b95d0ad53909d069a12086496903c7c928))
- Rename building file from pug to livereloader ([fd2d239](fd2d23982bed26ca8cee6e53548017f2f2777e03))
- Replace if to switch-case ([46be7da](46be7da7505d49c33af6f771f5385b136a0950ac))
- Improve readability of swithc-case operator ([714040c](714040c0bda6e4e5e6f2276bad0c01281e7c7c77))
- Rename cmd/s3-go-saver to cmd/cli ([4f8425b](4f8425b4981a96a0ab6460fc8844b4f88b816c8f))
- Switch-case: add \n for default help message ([2903fe0](2903fe0b58ebbc5f03126b992c87d5690b849d19))
- -time: split declaration of the start var from var assignment ([fca4070](fca4070d4e03046306a85c38e06bde622b5fab72))
- Move to package main ([b59e897](b59e897d77a8e7bfb717fe1c4983991a69f0e7b0))
- Replace key -fuzzy-search to -fuzzy ([68c8939](68c89391c5c7907bcee37f869c2d54f77e2e712d))
- Flag -download is []string ([5c967a2](5c967a22c2bb01c1a47278ae1351d7825139f6ef))
- Rename ListBucketOutput to BucketObjects ([e64ebe4](e64ebe46b5ebcf2cd6b2d64e66525f2c8a0dd0c7))
- Incapsulate code from flag -download-all to DownloadAllFiles() method ([110ea92](110ea9245883ffa5fbb24e77c92c3c622fe61591))

<!-- generated by git-cliff -->
