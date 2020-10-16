package main

import (
	"context"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"proxypool-go/util/iputil"
	"sync"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/shutdown", HttpShutdownHandler)

	server := &http.Server{
		Addr:           "0.0.0.0:3001",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server run at:")
	fmt.Printf("- Local:   http://localhost:%s \r\n", "0.0.0.0")
	fmt.Printf("- Network: http://%s:%s \r\n", iputil.GetLocalHost(), "3001")

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: ", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// 使用WaitGroup同步Goroutine
	wg := &sync.WaitGroup{}

	graceFullShutdown(server, wg)

	logger.Println("waiting for the remaining connections to finish...")
	// 等待已经关闭的信号
	wg.Wait()

	logger.Println("Server exiting")
}

func graceFullShutdown(server *http.Server, wg *sync.WaitGroup) {
	quit := make(chan os.Signal, 1)
	// 监听 Ctrl+C 信号
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-quit:
		wg.Add(1)
		// 使用context控制 server.Shutdown 的超时时间
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		errs := server.Shutdown(ctx)
		if errs != nil {
			logger.Info("Server Shutdown:", errs)
			fmt.Println("Server Shutdown:", errs)
		}
		wg.Done()
	}
}

// HttpShutdownHandler .
func HttpShutdownHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("Hello world!"))
	}
}
