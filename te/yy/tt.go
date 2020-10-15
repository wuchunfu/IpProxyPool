package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//sigs := make(chan os.Signal, 1)
	//done := make(chan bool)
	//
	//signal.Notify(sigs, syscall.SIGINT)
	//
	//go func() {
	//	sig := <-sigs
	//	fmt.Println("shutting down, caused by ", sig)
	//	close(done)
	//}()
	//
	//<-done
	//fmt.Println("Graceful shutdown.")

	mux := http.NewServeMux()
	mux.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "Hello world!")
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go server.ListenAndServe()

	var wg sync.WaitGroup
	done := make(chan struct{}, 1)

	graceFullShutdown(server, &wg, done)

	wg.Wait()
}

func graceFullShutdown(server *http.Server, wg *sync.WaitGroup, done chan<- struct{}) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		wg.Add(1)
		//使用context控制srv.Shutdown的超时时间
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		fmt.Println("Server Shutdown:")
		errs := server.Shutdown(ctx)
		if errs != nil {
			fmt.Println("Server Shutdown:", errs)
		}
		fmt.Println("Server Shutdown:---")
		defer wg.Done()
		close(done)
	}
}
