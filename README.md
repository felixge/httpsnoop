# httpsnoop

Package httpsnoop provides an easy way to capture http related metrics (i.e.
response time, bytes written, and http status code) from your application's
http.Handlers.

Doing this requires non-trivial wrapping of the http.ResponseWriter interface,
which is also exposed for users interested in a more low-level API.

[![GoDoc](https://godoc.org/github.com/felixge/httpsnoop?status.svg)](https://godoc.org/github.com/felixge/httpsnoop)
[![Build Status](https://travis-ci.org/felixge/httpsnoop.svg?branch=master)](https://travis-ci.org/felixge/httpsnoop)

## Usage Example

```go
// myH is your app's http handler, perhaps a http.ServeMux or similar.
var myH http.Handler
// wrappedH wraps myH in order to log every request.
wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	m := CaptureMetrics(myH, w, r)
	log.Printf(
		"%s %s (code=%d dt=%s written=%d)",
		r.Method,
		r.URL,
		m.Code,
		m.Duration,
		m.Written,
	)
})
http.ListenAndServe(":8080", wrappedH)
```
