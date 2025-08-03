package adapters

import (
	"fmt"
	"s3storage/internal/s3"
	"s3storage/pkg/tui"

	"github.com/charmbracelet/bubbles/list"
)

type S3ListItems struct {
	S3 *s3.S3Client
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

func (a *S3ListItems) ListBucket() *[]list.Item {
	var items = []list.Item{}

	listBuckets := *a.S3.ListBucket()
	for _, object := range listBuckets {
		if object.Size != 0 {
			items = append(items, tui.Item{
				Top:  object.Key,
				Desc: fmt.Sprintf("Size: %s", FormatBytes(object.Size)),
			})
		}
	}
	return &items
}
