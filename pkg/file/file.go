package file

import (
	"io"
	"os"
)

type File struct{}

func New() File {
	return File{}
}

func (f File) CreateFile(fileName string) (bool, error) {
	isExists, err := f.PathExists(fileName)
	if err != nil {
		return false, err
	}

	if !isExists {
		file, err := os.Create(fileName)
		if err != nil {
			return false, err
		}
		defer file.Close()
	}

	return true, nil
}

func (f File) CreateDir(dirName string) (bool, error) {
	isExists, err := f.PathExists(dirName)
	if err != nil {
		return false, err
	}

	if !isExists {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (f File) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func (f File) AppendToFile(fileName string, content string) error {
	fi, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fi.Close()

	// 查找文件末尾的偏移量
	// 从末尾的偏移量开始写入内容
	n, _ := fi.Seek(0, io.SeekEnd)
	_, err = fi.WriteAt([]byte(content), n)

	return err
}
