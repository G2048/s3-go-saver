package adapters

import (
	"errors"
	"fmt"
	"os"
	"s3-go-saver/pkg/s3"
	"s3-go-saver/pkg/tui"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

type S3ListItems struct {
	S3                *s3.S3Client
	tabsItems         tui.TabsItems
	DownloadDir       string
	downloadsItems    []tui.Item
	nondownloadsItems []tui.Item
}

const (
	Exist     = "✅"
	DontExist = "❌"
)
const RootTab = "/"

var mapExist = map[bool]string{true: Exist, false: DontExist}

// FormatBytes takes a byte size (int64) and returns a human-readable string.
func FormatBytes(b int64) string {
	if b == 0 {
		return "0 B"
	}
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div := int64(unit)
	exp := 0
	for absB := b; absB >= unit && exp < 8; absB /= unit { // Limiting to YB
		div *= unit
		exp++
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	return fmt.Sprintf("%.2f %s", float64(b)/float64(div/unit), units[exp])
}

func splitPath(key string) (string, string) {
	parts := strings.Split(key, "/")
	return parts[0], strings.Join(parts[1:], "/")
}
func checkExist(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return !errors.Is(err, os.ErrNotExist) && err == nil

}
func (a *S3ListItems) GetTabs() []tui.Tab {
	var tabs []tui.Tab
	var setTabs = make(map[tui.Tab]bool)

	if a.nondownloadsItems == nil && a.downloadsItems == nil {
		a.GetTabsItems()
	}
	for _, ti := range a.nondownloadsItems {
		setTabs[ti.Tab] = true
	}
	for _, ti := range a.downloadsItems {
		setTabs[ti.Tab] = true
	}
	for key, _ := range setTabs {
		tabs = append(tabs, key)
	}
	return tabs
}
func (a *S3ListItems) GetTabsItems() tui.TabsItems {
	var tabsItems = make(tui.TabsItems)

	listBuckets := *a.S3.ListBucket()
	for _, object := range listBuckets {
		if object.Size != 0 {
			tab, obj := splitPath(object.Key)
			if obj == "" {
				obj = tab
				tab = RootTab
			}
			exist := checkExist(a.DownloadDir + "/" + object.Key)
			desc := fmt.Sprintf("Size: %s; Downloaded: %s", FormatBytes(object.Size), mapExist[exist])
			ti := tui.Item{
				Top:      obj,
				Desc:     desc,
				Download: exist,
				Tab:      tui.Tab(tab),
			}
			if exist {
				a.downloadsItems = append(a.downloadsItems, ti)
			} else {
				a.nondownloadsItems = append(a.nondownloadsItems, ti)
			}
			if _, ok := tabsItems[ti.Tab]; !ok {
				tabsItems[ti.Tab] = []list.Item{ti}
			} else {
				tabsItems[ti.Tab] = append(tabsItems[ti.Tab], ti)
			}
		}
	}
	// a.tabsItems = tabsItems
	return tabsItems
}

func restoreFullPath(tab tui.Tab, file tui.Item) string {
	if tab == RootTab {
		return file.Top
	} else {
		return string(tab) + "/" + file.Top
	}
}

func (a *S3ListItems) DownloadItem(tab tui.Tab, file tui.Item) tui.Item {
	if file.Download {
		return file
	}
	fullPath := restoreFullPath(tab, file)
	err := a.S3.DownloadFile(fullPath, a.DownloadDir, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	file.Download = true
	file.Desc = strings.Replace(file.Desc, DontExist, Exist, 1)
	return file
}
func (a *S3ListItems) DeleteItem(tab tui.Tab, file tui.Item) tui.Item {
	if !file.Download {
		return file
	}

	fullPath := a.DownloadDir + "/" + restoreFullPath(tab, file)
	err := os.Remove(fullPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	file.Download = false
	file.Desc = strings.Replace(file.Desc, Exist, DontExist, 1)

	// Delete empty dir
	entries, err := os.ReadDir(a.DownloadDir + "/" + string(tab))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if len(entries) == 0 {
		if tab != "/" {
			os.Remove(a.DownloadDir + "/" + string(tab))
		}
	}
	return file
}
