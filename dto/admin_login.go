package dto

import (
	"time"

	"github.com/CoderTH/go_gateway/public"
	"github.com/gin-gonic/gin"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required,is_valid_username"` //管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`                   //密码
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token
}
