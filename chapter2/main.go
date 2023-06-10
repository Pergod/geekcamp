package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	rHeader := r.Header
	for k, v := range rHeader {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(200)
	io.WriteString(w, "ok")
	log.Printf("Host: %s , Http Code: %d", r.Host)
}
