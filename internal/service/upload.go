package service

import (
	"errors"
	"mime/multipart"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/file"
	"github.com/lughong/blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func (srv Service) UploadFile(t upload.FileType, header *multipart.FileHeader, f multipart.File) (*FileInfo, error) {
	dir := upload.GetSavePath()
	fileName := upload.GetFileName(header.Filename)
	dst := dir + "/" + fileName

	if upload.CheckMaxSize(t, f) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	if !upload.CheckContainExt(t, fileName) {
		return nil, errors.New("file suffix is not support.")
	}

	if upload.CheckPermission(fileName) {
		return nil, errors.New("insufficient file permissions.")
	}

	fi := file.New()
	if _, err := fi.CreateDir(dir); err != nil {
		return nil, err
	}

	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}

	accessURL := global.AppSetting.UploadserverUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessURL: accessURL}, nil
}
