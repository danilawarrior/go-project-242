package code

import (
	"os"
)

func GetPathSize(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		return getSizeOfDir(path)
	}

	return fileInfo.Size(), nil
}

func getSizeOfDir(path string) (int64, error) {

	data, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var sizeOfDir int64
	for _, entry := range data {
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		if !info.IsDir() {
			sizeOfDir += info.Size()
		}
	}
	return sizeOfDir, nil
}
