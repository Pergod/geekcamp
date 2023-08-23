package main

import (
	"fmt"
	"geekcamp/chapter10/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/latency", latency)
	mux.HandleFunc("/healthz", healthz)
	mux.Handle("/metrics", promhttp.Handler())

	srv := http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func latency(w http.ResponseWriter, request *http.Request) {
	// 随机时延
	startTime := time.Now()
	delay := rand.Int31n(2)
	time.Sleep(time.Second * time.Duration(delay))
	io.WriteString(w, "ok")
	duration := time.Since(startTime).Seconds()
	metrics.HttpRequestLatency.WithLabelValues("/example").Observe(duration)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 写入request header至response header
	for k, v := range r.Header {
		w.Header().Set(k, fmt.Sprintf("%s=%s\n", k, v))
	}
	// 获取环境变量version信息
	version := os.Getenv("VERSION")
	if len(version) > 0 {
		// 设置version至header
		w.Header().Set("VERSION", version)
	}
	// 设置200状态码
	w.WriteHeader(200)
	io.WriteString(w, "ok")
	// 访问日志记录host , http code
	log.Printf("Host: %s , Http Code: %d", r.Host, http.StatusOK)
}
