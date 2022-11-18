package httpsnoop

import "net/http"

type rw struct {
	w http.ResponseWriter
	h Hooks
	m *Metrics
}

func (w *rw) Metrics() *Metrics {
	return w.m
}

type Metricer interface {
	Metrics() *Metrics
}

func (w *rw) Unwrap() http.ResponseWriter {
	return w.w
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
