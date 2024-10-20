//go:build go1.8
// +build go1.8

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

// PushFunc is part of the http.Pusher interface.
type PushFunc func(target string, opts *http.PushOptions) error

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
	Push             func(PushFunc) PushFunc
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
// - http.Pusher
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
	_, i6 := w.(http.Pusher)
	switch {
	// combination 1/128
	case !i0 && !i1 && !i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
		}{rw, rw}
	// combination 2/128
	case !i0 && !i1 && !i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Pusher
		}{rw, rw, rw}
	// combination 3/128
	case !i0 && !i1 && !i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			fullDuplexEnabler
		}{rw, rw, rw}
	// combination 4/128
	case !i0 && !i1 && !i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 5/128
	case !i0 && !i1 && !i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
		}{rw, rw, rw}
	// combination 6/128
	case !i0 && !i1 && !i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 7/128
	case !i0 && !i1 && !i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 8/128
	case !i0 && !i1 && !i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 9/128
	case !i0 && !i1 && !i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 10/128
	case !i0 && !i1 && !i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 11/128
	case !i0 && !i1 && !i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 12/128
	case !i0 && !i1 && !i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 13/128
	case !i0 && !i1 && !i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw}
	// combination 14/128
	case !i0 && !i1 && !i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 15/128
	case !i0 && !i1 && !i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 16/128
	case !i0 && !i1 && !i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 17/128
	case !i0 && !i1 && i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
		}{rw, rw, rw}
	// combination 18/128
	case !i0 && !i1 && i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 19/128
	case !i0 && !i1 && i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 20/128
	case !i0 && !i1 && i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 21/128
	case !i0 && !i1 && i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw}
	// combination 22/128
	case !i0 && !i1 && i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 23/128
	case !i0 && !i1 && i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 24/128
	case !i0 && !i1 && i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 25/128
	case !i0 && !i1 && i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 26/128
	case !i0 && !i1 && i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 27/128
	case !i0 && !i1 && i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 28/128
	case !i0 && !i1 && i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 29/128
	case !i0 && !i1 && i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 30/128
	case !i0 && !i1 && i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 31/128
	case !i0 && !i1 && i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 32/128
	case !i0 && !i1 && i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 33/128
	case !i0 && i1 && !i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw, rw}
	// combination 34/128
	case !i0 && i1 && !i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 35/128
	case !i0 && i1 && !i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 36/128
	case !i0 && i1 && !i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 37/128
	case !i0 && i1 && !i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
		}{rw, rw, rw, rw}
	// combination 38/128
	case !i0 && i1 && !i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 39/128
	case !i0 && i1 && !i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 40/128
	case !i0 && i1 && !i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 41/128
	case !i0 && i1 && !i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 42/128
	case !i0 && i1 && !i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 43/128
	case !i0 && i1 && !i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 44/128
	case !i0 && i1 && !i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 45/128
	case !i0 && i1 && !i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 46/128
	case !i0 && i1 && !i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 47/128
	case !i0 && i1 && !i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 48/128
	case !i0 && i1 && !i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 49/128
	case !i0 && i1 && i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 50/128
	case !i0 && i1 && i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 51/128
	case !i0 && i1 && i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 52/128
	case !i0 && i1 && i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 53/128
	case !i0 && i1 && i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 54/128
	case !i0 && i1 && i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 55/128
	case !i0 && i1 && i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 56/128
	case !i0 && i1 && i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 57/128
	case !i0 && i1 && i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 58/128
	case !i0 && i1 && i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 59/128
	case !i0 && i1 && i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 60/128
	case !i0 && i1 && i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 61/128
	case !i0 && i1 && i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 62/128
	case !i0 && i1 && i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 63/128
	case !i0 && i1 && i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 64/128
	case !i0 && i1 && i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 65/128
	case i0 && !i1 && !i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
		}{rw, rw, rw}
	// combination 66/128
	case i0 && !i1 && !i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 67/128
	case i0 && !i1 && !i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			fullDuplexEnabler
		}{rw, rw, rw, rw}
	// combination 68/128
	case i0 && !i1 && !i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 69/128
	case i0 && !i1 && !i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
		}{rw, rw, rw, rw}
	// combination 70/128
	case i0 && !i1 && !i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 71/128
	case i0 && !i1 && !i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 72/128
	case i0 && !i1 && !i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 73/128
	case i0 && !i1 && !i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 74/128
	case i0 && !i1 && !i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 75/128
	case i0 && !i1 && !i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 76/128
	case i0 && !i1 && !i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 77/128
	case i0 && !i1 && !i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 78/128
	case i0 && !i1 && !i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 79/128
	case i0 && !i1 && !i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 80/128
	case i0 && !i1 && !i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 81/128
	case i0 && !i1 && i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 82/128
	case i0 && !i1 && i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 83/128
	case i0 && !i1 && i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 84/128
	case i0 && !i1 && i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 85/128
	case i0 && !i1 && i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 86/128
	case i0 && !i1 && i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 87/128
	case i0 && !i1 && i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 88/128
	case i0 && !i1 && i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 89/128
	case i0 && !i1 && i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 90/128
	case i0 && !i1 && i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 91/128
	case i0 && !i1 && i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 92/128
	case i0 && !i1 && i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 93/128
	case i0 && !i1 && i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 94/128
	case i0 && !i1 && i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 95/128
	case i0 && !i1 && i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 96/128
	case i0 && !i1 && i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 97/128
	case i0 && i1 && !i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw, rw}
	// combination 98/128
	case i0 && i1 && !i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 99/128
	case i0 && i1 && !i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw}
	// combination 100/128
	case i0 && i1 && !i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 101/128
	case i0 && i1 && !i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
		}{rw, rw, rw, rw, rw}
	// combination 102/128
	case i0 && i1 && !i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 103/128
	case i0 && i1 && !i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 104/128
	case i0 && i1 && !i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 105/128
	case i0 && i1 && !i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 106/128
	case i0 && i1 && !i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 107/128
	case i0 && i1 && !i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 108/128
	case i0 && i1 && !i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 109/128
	case i0 && i1 && !i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 110/128
	case i0 && i1 && !i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 111/128
	case i0 && i1 && !i2 && i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 112/128
	case i0 && i1 && !i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 113/128
	case i0 && i1 && i2 && !i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw, rw}
	// combination 114/128
	case i0 && i1 && i2 && !i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 115/128
	case i0 && i1 && i2 && !i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw}
	// combination 116/128
	case i0 && i1 && i2 && !i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 117/128
	case i0 && i1 && i2 && !i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
		}{rw, rw, rw, rw, rw, rw}
	// combination 118/128
	case i0 && i1 && i2 && !i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 119/128
	case i0 && i1 && i2 && !i3 && i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 120/128
	case i0 && i1 && i2 && !i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 121/128
	case i0 && i1 && i2 && i3 && !i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw, rw}
	// combination 122/128
	case i0 && i1 && i2 && i3 && !i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 123/128
	case i0 && i1 && i2 && i3 && !i4 && i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 124/128
	case i0 && i1 && i2 && i3 && !i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 125/128
	case i0 && i1 && i2 && i3 && i4 && !i5 && !i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
		}{rw, rw, rw, rw, rw, rw, rw}
	// combination 126/128
	case i0 && i1 && i2 && i3 && i4 && !i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw}
	// combination 127/128
	case i0 && i1 && i2 && i3 && i4 && i5 && !i6:
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
	// combination 128/128
	case i0 && i1 && i2 && i3 && i4 && i5 && i6:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			deadliner
			fullDuplexEnabler
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw, rw, rw}
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

func (w *rw) Push(target string, opts *http.PushOptions) error {
	f := w.w.(http.Pusher).Push
	if w.h.Push != nil {
		f = w.h.Push(f)
	}
	return f(target, opts)
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
