package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is home page"))
}

// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {

	g, ctx := errgroup.WithContext(context.Background())
	// https://golang.google.cn/doc/go1.17#vet
	// 注意自1.17开始这里的chan必须是带缓冲的
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan            // 此处没有系统信号时阻塞，后续代码不执行，有信号时后续代码执行
		signal.Stop(signalChan) // 显式停止监听系统信号
		close(signalChan)       // 显式关闭监听信号的通道
	}()

	mh := MyHandler{}
	http.Handle("/", mh)

	srv := &http.Server{
		Addr: ":8081",
	}

	// 主进程（主协程）阻塞channel，以便控制http-server退出后才退出主进程（主协程）
	// 就是一个简单的空结构体channel
	idleCloser := make(chan struct{})

	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			close(idleCloser)
			return err
		}
		return nil
	})

	g.Go(func() error {
		// 启动一个监听系统信号控制的channel
		<-signalChan

		// 超时context
		_, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("\nHttp服务暴力停止，一般是达到超时context的时间当前还有尚未完结的http请求：" + err.Error())
		} else {
			fmt.Println("\nHttp服务优雅停止")
		}

		// 关闭 主进程（主协程）阻塞channel
		close(idleCloser)

		return nil
	})

	g.Wait()

	fmt.Println("exit successful")
}
