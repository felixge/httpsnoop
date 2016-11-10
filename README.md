# httpsnoop

Package httpsnoop provides an easy way to capture http related metrics (i.e.
response time, bytes written, and http status code) from your application's
http.Handlers.

Doing this requires non-trivial wrapping of the http.ResponseWriter interface,
which is also exposed for users interested in a more low-level API.

[![GoDoc](https://godoc.org/github.com/felixge/httpsnoop?status.svg)](https://godoc.org/github.com/felixge/httpsnoop)
