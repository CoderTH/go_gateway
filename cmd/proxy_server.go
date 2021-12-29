package main

import (
	"github.com/CoderTH/go_gateway/http_proxy_router"
	"os"
	"os/signal"
	"syscall"

	"github.com/e421083458/golang_common/lib"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	go func() {
		http_proxy_router.HttpServerRun()
	}()
	go func() {
		http_proxy_router.HttpsServerRun()
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	http_proxy_router.HttpServerStop()
	http_proxy_router.HttpsServerStop()
}
