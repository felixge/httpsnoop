//go:build !go1.8
// +build !go1.8

// Code generated by "httpsnoop/codegen"; DO NOT EDIT.

package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"time"
)

// HeaderFunc is part of the http.ResponseWriter interface.
type HeaderFunc func() http.Header

// WriteHeaderFunc is part of the http.ResponseWriter interface.
type WriteHeaderFunc func(code int)

// WriteFunc is part of the http.ResponseWriter interface.
type WriteFunc func(b []byte) (int, error)

// FlushFunc is part of the http.Flusher interface.
type FlushFunc func()

// CloseNotifyFunc is part of the http.CloseNotifier interface.
type CloseNotifyFunc func() <-chan bool

// HijackFunc is part of the http.Hijacker interface.
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

// ReadFromFunc is part of the io.ReaderFrom interface.
type ReadFromFunc func(src io.Reader) (int64, error)

// SetReadDeadlineFunc is part of the deadliner interface.
type SetReadDeadlineFunc func(deadline time.Time) error

// SetWriteDeadlineFunc is part of the deadliner interface.
type SetWriteDeadlineFunc func(deadline time.Time) error

// EnableFullDuplexFunc is part of the fullDuplexEnabler interface.
type EnableFullDuplexFunc func() error

// Hooks defines a set of method interceptors for methods included in
// http.ResponseWriter as well as some others. You can think of them as
// middleware for the function calls they target. See Wrap for more details.
type Hooks struct {
	Header           func(HeaderFunc) HeaderFunc
	WriteHeader      func(WriteHeaderFunc) WriteHeaderFunc
	Write            func(WriteFunc) WriteFunc
	Flush            func(FlushFunc) FlushFunc
	CloseNotify      func(CloseNotifyFunc) CloseNotifyFunc
	Hijack           func(HijackFunc) HijackFunc
	ReadFrom         func(ReadFromFunc) ReadFromFunc
	SetReadDeadline  func(SetReadDeadlineFunc) SetReadDeadlineFunc
	SetWriteDeadline func(SetWriteDeadlineFunc) SetWriteDeadlineFunc
	EnableFullDuplex func(EnableFullDuplexFunc) EnableFullDuplexFunc
}

// Wrap returns a wrapped version of w that provides the exact same interface
// as w. Specifically if w implements any combination of:
//
// - http.Flusher
// - http.CloseNotifier
// - http.Hijacker
// - io.ReaderFrom
// - deadliner
// - fullDuplexEnabler
//
// The wrapped version will implement the exact same combination. If no hooks
// are set, the wrapped version also behaves exactly as w. Hooks targeting
// methods not supported by w are ignored. Any other hooks will intercept the
// method they target and may modify the call's arguments and/or return values.
// The CaptureMetrics implementation serves as a working example for how the
// hooks can be used.
func Wrap(w http.ResponseWriter, hooks Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: hooks}
	_, i0 := w.(http.Flusher)
	_, i1 := w.(http.CloseNotifier)
	_, i2 := w.(http.Hijacker)
	_, i3 := w.(io.ReaderFrom)
	_, i4 := w.(deadliner)
	_, i5 := w.(fullDuplexEnabler)
	switch {
	// combination 1/64
	case !i0 && !i1 && !i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
		}{rw, rw}
	// combination 2/64
	case !i0 && !i1 && !i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			fullDuplexEnabler
		}{rw, rw, rw}
	// combination 3/64
	case !i0 && !i1 && !i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
		}{rw, rw, rw}
	// combination 4/64
	case !i0 && !i1 && !i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 5/64
	case !i0 && !i1 && !i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 6/64
	case !i0 && !i1 && !i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 7/64
	case !i0 && !i1 && !i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw}
	// combination 8/64
	case !i0 && !i1 && !i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 9/64
	case !i0 && !i1 && i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
		}{rw, rw, rw}
	// combination 10/64
	case !i0 && !i1 && i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 11/64
	case !i0 && !i1 && i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw}
	// combination 12/64
	case !i0 && !i1 && i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 13/64
	case !i0 && !i1 && i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 14/64
	case !i0 && !i1 && i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 15/64
	case !i0 && !i1 && i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 16/64
	case !i0 && !i1 && i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 17/64
	case !i0 && i1 && !i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw, rw}
	// combination 18/64
	case !i0 && i1 && !i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 19/64
	case !i0 && i1 && !i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
		}{rw, rw, rw, rw}
	// combination 20/64
	case !i0 && i1 && !i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 21/64
	case !i0 && i1 && !i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 22/64
	case !i0 && i1 && !i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 23/64
	case !i0 && i1 && !i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 24/64
	case !i0 && i1 && !i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 25/64
	case !i0 && i1 && i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 26/64
	case !i0 && i1 && i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 27/64
	case !i0 && i1 && i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 28/64
	case !i0 && i1 && i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 29/64
	case !i0 && i1 && i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 30/64
	case !i0 && i1 && i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 31/64
	case !i0 && i1 && i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 32/64
	case !i0 && i1 && i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 33/64
	case i0 && !i1 && !i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
		}{rw, rw, rw}
	// combination 34/64
	case i0 && !i1 && !i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 35/64
	case i0 && !i1 && !i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
		}{rw, rw, rw, rw}
	// combination 36/64
	case i0 && !i1 && !i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 37/64
	case i0 && !i1 && !i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 38/64
	case i0 && !i1 && !i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 39/64
	case i0 && !i1 && !i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 40/64
	case i0 && !i1 && !i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 41/64
	case i0 && !i1 && i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 42/64
	case i0 && !i1 && i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 43/64
	case i0 && !i1 && i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 44/64
	case i0 && !i1 && i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 45/64
	case i0 && !i1 && i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 46/64
	case i0 && !i1 && i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 47/64
	case i0 && !i1 && i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 48/64
	case i0 && !i1 && i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 49/64
	case i0 && i1 && !i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw, rw}
	// combination 50/64
	case i0 && i1 && !i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 51/64
	case i0 && i1 && !i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 52/64
	case i0 && i1 && !i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 53/64
	case i0 && i1 && !i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 54/64
	case i0 && i1 && !i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 55/64
	case i0 && i1 && !i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 56/64
	case i0 && i1 && !i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 57/64
	case i0 && i1 && i2 && !i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw, rw}
	// combination 58/64
	case i0 && i1 && i2 && !i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 59/64
	case i0 && i1 && i2 && !i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 60/64
	case i0 && i1 && i2 && !i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 61/64
	case i0 && i1 && i2 && i3 && !i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw, rw}
	// combination 62/64
	case i0 && i1 && i2 && i3 && !i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 63/64
	case i0 && i1 && i2 && i3 && i4 && !i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 64/64
	case i0 && i1 && i2 && i3 && i4 && i5:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	}
	panic("unreachable")
}

type rw struct {
	w http.ResponseWriter
	h Hooks
}

func (w *rw) Unwrap() http.ResponseWriter {
	return w.w
}

func (w *rw) Header() http.Header {
	f := w.w.(http.ResponseWriter).Header
	if w.h.Header != nil {
		f = w.h.Header(f)
	}
	return f()
}

func (w *rw) WriteHeader(code int) {
	f := w.w.(http.ResponseWriter).WriteHeader
	if w.h.WriteHeader != nil {
		f = w.h.WriteHeader(f)
	}
	f(code)
}

func (w *rw) Write(b []byte) (int, error) {
	f := w.w.(http.ResponseWriter).Write
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

func (w *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.w.(http.Hijacker).Hijack
	if w.h.Hijack != nil {
		f = w.h.Hijack(f)
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

func (w *rw) SetReadDeadline(deadline time.Time) error {
	f := w.w.(deadliner).SetReadDeadline
	if w.h.SetReadDeadline != nil {
		f = w.h.SetReadDeadline(f)
	}
	return f(deadline)
}

func (w *rw) SetWriteDeadline(deadline time.Time) error {
	f := w.w.(deadliner).SetWriteDeadline
	if w.h.SetWriteDeadline != nil {
		f = w.h.SetWriteDeadline(f)
	}
	return f(deadline)
}

func (w *rw) EnableFullDuplex() error {
	f := w.w.(fullDuplexEnabler).EnableFullDuplex
	if w.h.EnableFullDuplex != nil {
		f = w.h.EnableFullDuplex(f)
	}
	return f()
}

type Unwrapper interface {
	Unwrap() http.ResponseWriter
}

// Unwrap returns the underlying http.ResponseWriter from within zero or more
// layers of httpsnoop wrappers.
func Unwrap(w http.ResponseWriter) http.ResponseWriter {
	if rw, ok := w.(Unwrapper); ok {
		// recurse until rw.Unwrap() returns a non-Unwrapper
		return Unwrap(rw.Unwrap())
	} else {
		return w
	}
}
