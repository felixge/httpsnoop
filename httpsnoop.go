package httpsnoop

import (
	"net/http"
	"time"
)

type HeaderFunc func() http.Header
type WriteFunc func([]byte) (int, error)
type WriteHeaderFunc func(int)
type FlushFunc func()
type CloseNotifyFunc func() <-chan bool

type Hooks struct {
	Header      func(next HeaderFunc) HeaderFunc
	Write       func(next WriteFunc) WriteFunc
	WriteHeader func(next WriteHeaderFunc) WriteHeaderFunc
	Flush       func(next FlushFunc) FlushFunc
	CloseNotify func(next CloseNotifyFunc) CloseNotifyFunc
}

func SnoopResponseWriter(w http.ResponseWriter, h Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: h}
	_, fOk := w.(http.Flusher)
	_, cnOk := w.(http.CloseNotifier)
	if fOk && cnOk {
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw}
	} else if fOk {
		return struct {
			http.ResponseWriter
			http.Flusher
		}{rw, rw}
	} else if cnOk {
		return struct {
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw}
	} else {
		return struct {
			http.ResponseWriter
		}{rw}
	}
}

type rw struct {
	w http.ResponseWriter
	h Hooks
}

func (w *rw) Header() http.Header {
	fn := w.w.Header
	if w.h.Header != nil {
		fn = w.h.Header(fn)
	}
	return fn()
}

func (w *rw) WriteHeader(code int) {
	fn := w.w.WriteHeader
	if w.h.WriteHeader != nil {
		fn = w.h.WriteHeader(fn)
	}
	fn(code)
}

func (w *rw) Write(b []byte) (int, error) {
	f := w.w.Write
	if w.h.Write != nil {
		f = w.h.Write(f)
	}
	return f(b)
}

func (w *rw) Flush() {
	f := w.w.(http.Flusher).Flush
	if w.h.Flush != nil {
		f = w.h.Flush(f)
	}
	f()
}

func (w *rw) CloseNotify() <-chan bool {
	f := w.w.(http.CloseNotifier).CloseNotify
	if w.h.CloseNotify != nil {
		f = w.h.CloseNotify(f)
	}
	return f()
}

type Metrics struct {
	Code     int
	Duration time.Duration
	Written  int64
}

func SnoopMetrics(hnd http.Handler, w http.ResponseWriter, r *http.Request) Metrics {
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
		hnd.ServeHTTP(SnoopResponseWriter(w, hooks), r)
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
