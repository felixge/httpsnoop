package httpsnoop

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkBaseline(b *testing.B) {
	benchmark(b, 0)
}

func BenchmarkCaptureMetrics(b *testing.B) {
	benchmark(b, 1)
}

func BenchmarkCaptureMetricsTwice(b *testing.B) {
	benchmark(b, 2)
}

func benchmark(b *testing.B, wrappings int) {
	dummyH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	h := dummyH
	for x := 0; x < wrappings; x++ {
		hCopy := h
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			CaptureMetrics(hCopy, w, r)
		})
	}

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	resp := httptest.NewRecorder() // ok to reuse; we're not writing anything to it

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(resp, req)
	}
}
