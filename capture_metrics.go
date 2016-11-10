package httpsnoop

import (
	"net/http"
	"time"
)

type Metrics struct {
	Code     int
	Duration time.Duration
	Written  int64
}

func CaptureMetrics(hnd http.Handler, w http.ResponseWriter, r *http.Request) Metrics {
	var (
		start            = time.Now()
		m                = Metrics{Code: http.StatusOK}
		writeHeaderCount int
		updates          = make(chan func())
		done             = make(chan struct{})
		hooks            = Hooks{
			WriteHeader: func(next WriteHeaderFunc) WriteHeaderFunc {
				return func(code int) {
					updates <- func() {
						if writeHeaderCount == 0 {
							m.Code = code
							writeHeaderCount++
						}
						next(code)
					}
				}
			},

			Write: func(next WriteFunc) WriteFunc {
				return func(p []byte) (int, error) {
					n, err := next(p)
					updates <- func() { m.Written += int64(n) }
					return n, err
				}
			},
		}
	)

	go func() {
		hnd.ServeHTTP(Wrap(w, hooks), r)
		close(done)
	}()

	for {
		select {
		case update := <-updates:
			update()
		case <-done:
			m.Duration = time.Since(start)
			return m
		}
	}
}
