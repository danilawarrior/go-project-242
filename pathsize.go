package code

import (
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, allFlag bool, isRecursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		return getSizeOfDir(path, allFlag, isRecursive)
	}

	return fileInfo.Size(), nil
}

func getSizeOfDir(path string, allFlag bool, isRecursive bool) (int64, error) {

	data, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var sizeOfDir int64
	for _, entry := range data {
		if !allFlag && strings.HasPrefix(entry.Name(), ".") { //если флаг не задан и файл скрытый, то его не считаем.
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		if !info.IsDir() {
			sizeOfDir += info.Size()
		} else if isRecursive {
			sizeOfDirInside, err := getSizeOfDir(filepath.Join(path, info.Name()), allFlag, isRecursive)
			if err != nil {
				return 0, err
			}
			sizeOfDir += sizeOfDirInside
		}
	}
	return sizeOfDir, nil
}
