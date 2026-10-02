package main

import (
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello\n"))
	})
	wrapped := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, w, r)
		log.Printf("%s %s code=%d duration=%s written=%d", r.Method, r.URL, m.Code, m.Duration, m.Written)
	})
	log.Fatal(http.ListenAndServe(":8080", wrapped))
}
