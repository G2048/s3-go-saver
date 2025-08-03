package adapters

import (
	"fmt"
	"s3storage/internal/s3"
	"s3storage/pkg/tui"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

type S3ListItems struct {
	S3          *s3.S3Client
	tabsItems   tui.TabsItems
	DownloadDir string
}

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

func (a S3ListItems) GetTabs() []tui.Tab {
	var tabs []tui.Tab
	if a.tabsItems == nil {
		a.tabsItems = a.GetTabsItems()
	}
	for tab, _ := range a.tabsItems {
		tabs = append(tabs, tab)
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
				tab = "/"
			}
			ti := tui.Item{
				Top:  obj,
				Desc: fmt.Sprintf("Size: %s", FormatBytes(object.Size)),
			}
			if _, ok := tabsItems[tui.Tab(tab)]; !ok {
				tabsItems[tui.Tab(tab)] = []list.Item{ti}
			} else {
				tabsItems[tui.Tab(tab)] = append(tabsItems[tui.Tab(tab)], ti)
			}
		}
	}
	a.tabsItems = tabsItems
	return tabsItems
}
func (a *S3ListItems) DownloadItems(file string) {
	a.S3.DownloadFile(file, a.DownloadDir)
}
