package httpsnoop

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCaptureMetrics(t *testing.T) {
	tests := []struct {
		Handler      http.Handler
		WantDuration time.Duration
		WantWritten  int64
		WantCode     int
	}{
		{
			Handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			WantCode: http.StatusOK,
		},
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("foo"))
				w.Write([]byte("bar"))
				time.Sleep(25 * time.Millisecond)
			}),
			WantCode:     http.StatusBadRequest,
			WantWritten:  6,
			WantDuration: 25 * time.Millisecond,
		},
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("foo"))
				w.WriteHeader(http.StatusNotFound)
			}),
			WantCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		func() {
			ch := make(chan Metrics, 1)
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ch <- CaptureMetrics(test.Handler, w, r)
			})
			s := httptest.NewServer(h)
			defer s.Close()
			res, err := http.Get(s.URL)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			m := <-ch
			if m.Code != test.WantCode {
				t.Errorf("got=%d want=%d", m.Code, test.WantCode)
			} else if m.Duration < test.WantDuration {
				t.Errorf("got=%s want=%s", m.Duration, test.WantDuration)
			} else if m.Written < test.WantWritten {
				t.Errorf("got=%d want=%d", m.Written, test.WantWritten)
			}
		}()
	}
}
