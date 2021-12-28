package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/e421083458/golang_common/lib"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	fmt.Println("start proxy server")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
