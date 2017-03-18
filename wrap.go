//go:generate bash -c "go run codegen/main.go > wrap_gen.go"

package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// HeaderFunc is part of the http.ResponseWriter interface.
type HeaderFunc func() http.Header

// WriteHeaderFunc is part of the http.ResponseWriter interface.
type WriteHeaderFunc func(int)

// WriteFunc is part of the http.ResponseWriter interface.
type WriteFunc func([]byte) (int, error)

// FlushFunc is part of the http.Flusher interface.
type FlushFunc func()

// CloseNotifyFunc is part of the http.CloseNotifier interface.
type CloseNotifyFunc func() <-chan bool

// HijackFunc is part of the http.Hijacker interface.
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

// ReadFromFunc is part of the io.ReaderFrom interface.
type ReadFromFunc func(src io.Reader) (int64, error)

// Hooks defines a set of method interceptors for methods included in
// http.ResponseWriter as well as some others. You can think of them as
// middleware for the function calls they target. See Wrap for more details.
type Hooks struct {
	Header      func(HeaderFunc) HeaderFunc
	Write       func(WriteFunc) WriteFunc
	WriteHeader func(WriteHeaderFunc) WriteHeaderFunc
	Flush       func(FlushFunc) FlushFunc
	CloseNotify func(CloseNotifyFunc) CloseNotifyFunc
	ReadFrom    func(ReadFromFunc) ReadFromFunc
	Hijack      func(HijackFunc) HijackFunc
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
