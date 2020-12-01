package app

import (
	"strings"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/file"
)

func CheckUploadFileMaxSize(fileSize uint8) bool {
	if fileSize > global.AppSetting.UploadFileMaxSize {
		return false
	}

	return true
}

func CheckUploadFileExt(fileExt string) bool {
	for _, ext := range global.AppSetting.UploadFileExt {
		if strings.ToLower(fileExt) == ext {
			return true
		}
	}

	return false
}

func CheckUploadSavePath() (bool, error) {
	f := file.New()

	uploadDir := global.RootDir + "/" + global.AppSetting.UploadFileSavePath
	if _, err := f.CreateDir(uploadDir); err != nil {
		return false, err
	}

	return true, nil
}

func CreateUploadFileName() {}
