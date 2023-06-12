package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
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
