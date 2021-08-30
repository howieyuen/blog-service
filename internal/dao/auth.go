package dao

import (
	"github.com/go-programming-tour/blog-service/internal/model"
)

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := &model.Auth{
		Model: &model.Model{
			IsDel: 0,
		},
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
