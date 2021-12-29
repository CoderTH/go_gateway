package http_proxy_middleware

import (
	"errors"
	"github.com/CoderTH/go_gateway/dao"
	"github.com/CoderTH/go_gateway/middleware"
	"github.com/CoderTH/go_gateway/public"
	"github.com/gin-gonic/gin"
	"strings"
)

//当设置为前缀匹配并且开启了strip_uri时，请求下游时会删除前缀
//例如：本来127.0.0.1:8080/test/abc 会请求下游：127.0.0.1:2003/test/abc
//开启后 127.0.0.1:8080/test/abc 会请求下游：127.0.0.1:2003/abc
func HttpStripURIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		if serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HTTPRule.NeedStripUri == 1 {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, serviceDetail.HTTPRule.Rule, "", 1)
		}
		c.Next()
	}
}
