package file

import "os"

const (
	defaultDirectoryPermission = 755
)

type Options struct {
	Path string
}

func CreateDir(filePath string) error {
	if !FileExists(filePath) {
		err := os.MkdirAll(filePath, defaultDirectoryPermission)
		return err
	}
	return nil
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
