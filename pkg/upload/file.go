package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(dst string) string {
	ext := GetFileExt(dst)
	fileName := strings.TrimSuffix(dst, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.RootDir + "/" + global.AppSetting.UploadSavePath
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
				return true
			}
		}
	}

	return false
}

func CheckPermission(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsPermission(err)
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
