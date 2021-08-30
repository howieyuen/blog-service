package service

import (
	"errors"
	"mime/multipart"
	"os"
	
	"github.com/go-programming-tour/blog-service/global"
	"github.com/go-programming-tour/blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(
	fileType upload.FileType,
	file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo, error) {
	
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceed max file size")
	}
	
	savePath := upload.GetSavePath()
	if upload.CheckSavePath(savePath) {
		if err := upload.CreateSavePath(savePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save path")
		}
	}
	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permissions")
	}
	
	dst := savePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerURL + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
