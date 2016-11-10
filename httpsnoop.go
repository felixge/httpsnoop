package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"time"
)

type HeaderFunc func() http.Header
type WriteFunc func([]byte) (int, error)
type WriteHeaderFunc func(int)
type FlushFunc func()
type CloseNotifyFunc func() <-chan bool
type ReadFromFunc func(src io.Reader) (int64, error)
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

type Hooks struct {
	Header      func(HeaderFunc) HeaderFunc
	Write       func(WriteFunc) WriteFunc
	WriteHeader func(WriteHeaderFunc) WriteHeaderFunc
	Flush       func(FlushFunc) FlushFunc
	CloseNotify func(CloseNotifyFunc) CloseNotifyFunc
	ReadFrom    func(ReadFromFunc) ReadFromFunc
	Hijack      func(HijackFunc) HijackFunc
}

func Wrap(w http.ResponseWriter, hooks Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: hooks}
	_, h := w.(http.Hijacker)
	_, f := w.(http.Flusher)
	_, cn := w.(http.CloseNotifier)
	_, rf := w.(io.ReaderFrom)
	switch {
	case h && f && cn && rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	case h && f && cn && !rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw, rw}
	case h && f && !cn && rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw, rw}
	case h && f && !cn && !rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.Flusher
		}{rw, rw, rw}
	case h && !f && cn && rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	case h && !f && cn && !rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			http.CloseNotifier
		}{rw, rw, rw}
	case h && !f && !cn && rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw}
	case h && !f && !cn && !rf:
		return struct {
			http.ResponseWriter
			http.Hijacker
		}{rw, rw}
	case !h && f && cn && rf:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	case !h && f && cn && !rf:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw}
	case !h && f && !cn && rf:
		return struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw}
	case !h && f && !cn && !rf:
		return struct {
			http.ResponseWriter
			http.Flusher
		}{rw, rw}
	case !h && !f && cn && rf:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw}
	case !h && !f && cn && !rf:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw}
	case !h && !f && !cn && rf:
		return struct {
			http.ResponseWriter
			io.ReaderFrom
		}{rw, rw}
	case !h && !f && !cn && !rf:
		return struct {
			http.ResponseWriter
		}{rw}
	}
	panic("unreachable")
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

func (w *rw) ReadFrom(src io.Reader) (int64, error) {
	f := w.w.(io.ReaderFrom).ReadFrom
	if w.h.ReadFrom != nil {
		f = w.h.ReadFrom(f)
	}
	return f(src)
}

func (w *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.w.(http.Hijacker).Hijack
	if w.h.Hijack != nil {
		f = w.h.Hijack(f)
	}
	return f()
}

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
