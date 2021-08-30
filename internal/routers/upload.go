package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour/blog-service/global"
	"github.com/go-programming-tour/blog-service/internal/service"
	"github.com/go-programming-tour/blog-service/pkg/app"
	"github.com/go-programming-tour/blog-service/pkg/convert"
	"github.com/go-programming-tour/blog-service/pkg/errcode"
	"github.com/go-programming-tour/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile error: %v", err)
		response.ToErrorResponse(errcode.UploadFileFailed)
		return
	}
	response.ToResponse(fileInfo)
}