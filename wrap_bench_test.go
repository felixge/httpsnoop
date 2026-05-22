package httpsnoop

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkWrap(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	resp := httptest.NewRecorder()

	b.ReportAllocs()
	b.StartTimer()
	for b.Loop() {
		wrapped := Wrap(resp, Hooks{})
		if wrapped == nil {
			b.Fatal()
		}
	}
}

func BenchmarkWrappedWriteWithHook(b *testing.B) {
	var total int64
	w := httptest.NewRecorder()
	ww := Wrap(w, Hooks{
		Write: func(next WriteFunc) WriteFunc {
			return func(p []byte) (int, error) {
				n, err := next(p)
				total += int64(n)
				return n, err
			}
		},
	})
	payload := []byte("hello world")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ww.Write(payload)
		if err != nil {
			b.Fatal()
		}
	}
}

func BenchmarkWrappedWriteNoHook(b *testing.B) {
	w := httptest.NewRecorder()
	ww := Wrap(w, Hooks{})
	payload := []byte("hello world")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ww.Write(payload)
		if err != nil {
			b.Fatal()
		}
	}
}
