package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	var size int64

	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		size, err = getSizeOfDir(path, all, recursive)
	} else {
		size = fileInfo.Size()
	}
	if err != nil {
		return "", err
	}
	if human {
		return sizeToHuman(size), nil
	}

	return fmt.Sprintf("%d", size), nil
}

func getSizeOfDir(path string, all bool, recursive bool) (int64, error) {

	data, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var sizeOfDir int64
	for _, entry := range data {
		if !all && strings.HasPrefix(entry.Name(), ".") { //если флаг не задан и файл скрытый, то его не считаем.
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		if !info.IsDir() {
			sizeOfDir += info.Size()
		} else if recursive {
			sizeOfDirInside, err := getSizeOfDir(filepath.Join(path, info.Name()), all, recursive)
			if err != nil {
				return 0, err
			}
			sizeOfDir += sizeOfDirInside
		}
	}
	return sizeOfDir, nil
}

func sizeToHuman(size int64) string {
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
