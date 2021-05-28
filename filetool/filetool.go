package filetool

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// MkDirAll 递归创建文件夹
func MkDirAll(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// isExist判断所给路径是否存在(true:存在 false:No)
func isExist(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// CreateTempFile 创建临时文件
func CreateTempFile(dir string) (*os.File, error) {
	filePath := path.Join(os.TempDir(), dir)
	f, err := ioutil.TempFile(filePath, "export")
	return f, err
}

// CreateTeamDir 创建临时文件夹 ./Temp
func CreateTeamDir(dir string, removeOld bool) (string, error) {
	filePath := path.Join(os.TempDir(), dir)
	if removeOld {
		if err := os.RemoveAll(filePath); err != nil {
			return "", err
		}
	}
	return filePath, MkDirAll(filePath)
}

// RemoveFileAfter 在多长时间后移除文件
func RemoveFileAfter(filePath string, duration time.Duration) {
	time.AfterFunc(duration, func() {
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("RemoveFileAfter: %s", err.Error())
		}
	})
}
