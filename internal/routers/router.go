package routers

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour/blog-service/docs"
	"github.com/go-programming-tour/blog-service/global"
	"github.com/go-programming-tour/blog-service/internal/middleware"
	"github.com/go-programming-tour/blog-service/internal/routers/api"
	v1 "github.com/go-programming-tour/blog-service/internal/routers/api/v1"
	"github.com/go-programming-tour/blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// 中间件
	if global.ServerSetting.RunMode == "debug" {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	} else {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiter))
	r.Use(middleware.ContextTimeout(global.ServerSetting.ContextTimeout))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// 文件上传
	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	
	// 身份验证
	r.POST("/auth", api.GetAuth)
	
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)
		
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}

var methodLimiter = limiter.NewMethodLimiter().AddBucket(limiter.BucketRule{
	Key:      "/auth",
	Interval: time.Second,
	Capacity: 10,
	Quantum:  10,
})
