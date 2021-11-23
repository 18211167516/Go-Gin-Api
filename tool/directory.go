package tool

import (
	"errors"
	"go-api/global"
	"os"

	"go.uber.org/zap"
)

//@author: [18211167516](https://github.com/18211167516)
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [18211167516](https://github.com/18211167516)
//@function: CreateFile
//@description: 创建文件
//@param: filename string
//@return: os.File, error
func CreateFile(filename string) (*os.File, error) {
	exist, err := PathExists(filename)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("file already exist")
	} else {
		return os.Create(filename)
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
