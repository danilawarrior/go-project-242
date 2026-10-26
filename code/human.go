package code

import "fmt"

func SizeToHuman(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	if size < 1000 {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	sizeFloat := float64(size)
	unitIndex := 0
	for sizeFloat >= 1000 && unitIndex < len(units)-1 {
		sizeFloat /= 1000
		unitIndex++
	}
	return fmt.Sprintf("%.1f%s", sizeFloat, units[unitIndex])
}
