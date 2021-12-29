package http_proxy_middleware

import (
	"fmt"
	"github.com/CoderTH/go_gateway/dao"
	"github.com/CoderTH/go_gateway/middleware"
	"github.com/CoderTH/go_gateway/public"
	"github.com/gin-gonic/gin"
)

//匹配接入方式 基于请求信息
func HttpAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		fmt.Println("matched service :", public.ObjToJson(service))
		c.Set("service", service)
		c.Next()
	}
}
