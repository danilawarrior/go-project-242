package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return getSizeOfDir(path)
	}

	return fmt.Sprintf("%dB", fileInfo.Size()), nil
}

func getSizeOfDir(path string) (string, error) {

	data, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var sizeOfDir int64
	for _, entry := range data {
		info, err := entry.Info()
		if err != nil {
			return "", err
		}
		if !info.IsDir() {
			sizeOfDir += info.Size()
		}
	}
	return fmt.Sprintf("%dB", sizeOfDir), nil
}
