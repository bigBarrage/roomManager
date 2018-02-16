package run

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

		_ = <-sig
		fmt.Fprintln(os.Stdout, "正在关闭进程")
		//执行关闭钱需要处理的功能
		os.Exit(1)
	}()
}
