package httpsnoop

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkBaseline(b *testing.B) {
	benchmark(b, false)
}

func BenchmarkCaptureMetrics(b *testing.B) {
	benchmark(b, true)
}

func BenchmarkWrap(b *testing.B) {
	b.StopTimer()
	doneCh := make(chan struct{}, 1)
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Wrap(w, Hooks{})
		}
		doneCh <- struct{}{}
	})
	s := httptest.NewServer(h)
	defer s.Close()
	if _, err := http.Get(s.URL); err != nil {
		b.Fatal(err)
	}
	<-doneCh
}

func benchmark(b *testing.B, captureMetrics bool) {
	b.StopTimer()
	dummyH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	h := dummyH
	if captureMetrics {
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			CaptureMetrics(dummyH, w, r)
		})
	}
	s := httptest.NewServer(h)
	defer s.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := http.Get(s.URL)
		if err != nil {
			b.Fatal(err)
		}
	}
}
