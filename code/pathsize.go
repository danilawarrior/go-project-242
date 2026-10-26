package code

import (
	"os"
	"strings"
)

func GetPathSize(path string, allFlag bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		return getSizeOfDir(path, allFlag)
	}

	return fileInfo.Size(), nil
}

func getSizeOfDir(path string, allFlag bool) (int64, error) {

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
		}
	}
	return sizeOfDir, nil
}
