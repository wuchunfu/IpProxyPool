package api

import (
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"proxypool-go/middleware/storage"
	"proxypool-go/models/configModel"
	"proxypool-go/util/iputil"
	"sync"
	"syscall"
	"time"
)

// Run for request
func Run(setting *configModel.System) {

	mux := http.NewServeMux()
	mux.HandleFunc("/http", ProxyHttpHandler)
	mux.HandleFunc("/https", ProxyHttpsHandler)

	server := &http.Server{
		Addr:           setting.HttpAddr + ":" + setting.HttpPort,
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server run at:")
	fmt.Printf("-  Local:   http://localhost:%s \r\n", setting.HttpPort)
	fmt.Printf("-  Network: http://%s:%s \r\n", iputil.GetLocalHost(), setting.HttpPort)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("listen: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	server.SetKeepAlivesEnabled(false)
	errs := server.Shutdown(ctx)
	if errs != nil {
		logger.Info("Server Shutdown:", errs)
		fmt.Println("Server Shutdown:", errs)
	}

	//// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//// 使用WaitGroup同步Goroutine
	//wg := &sync.WaitGroup{}
	//
	//graceFullShutdown1(server, wg)
	//
	//fmt.Println("waiting for the remaining connections to finish...")
	//// 等待已经关闭的信号
	//wg.Wait()

	logger.Println("Server exiting")

	//fmt.Println("Starting server", setting.HttpAddr+":"+setting.HttpPort)
	//http.ListenAndServe(setting.HttpAddr+":"+setting.HttpPort, mux)
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

// ProxyHttpHandler .
func ProxyHttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyHttpRandom())
		if err != nil {
			return
		}
		w.Write(b)
	}
}

// ProxyHttpsHandler .
func ProxyHttpsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyHttpsRandom("https"))
		if err != nil {
			return
		}
		w.Write(b)
	}
}
