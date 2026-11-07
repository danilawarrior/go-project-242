package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

func getSize(path string, recursive, all bool) (int64, error) {

	//если path файл то вернуть его size
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	data, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var sizeOfDir int64
	for _, entry := range data {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() && !recursive {
			continue
		}

		size, err := getSize(entryPath, recursive, all)
		if err != nil {
			return 0, err
		}

		sizeOfDir += size
	}
	return sizeOfDir, nil

}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}

	sizeFloat := float64(size)
	unitIndex := 0

	for sizeFloat >= 1024 && unitIndex < len(units)-1 {
		sizeFloat /= 1024
		unitIndex++
	}

	return fmt.Sprintf("%.1f%s", sizeFloat, units[unitIndex])
}
