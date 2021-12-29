package http_proxy_router

import (
	"github.com/CoderTH/go_gateway/http_proxy_middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Use(
		http_proxy_middleware.HttpAccessModeMiddleware(),
		http_proxy_middleware.HttpHeaderTransferMiddleware(),
		http_proxy_middleware.HttpStripURIMiddleware(),
		http_proxy_middleware.HttpUrlRewriteMiddleware(),
		http_proxy_middleware.HTTPReverseProxyMiddleware(),
	)
	return router
}
