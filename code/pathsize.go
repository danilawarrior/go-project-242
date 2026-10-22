package code

import "os"

func GetPathSize(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}
