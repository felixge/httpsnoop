package httpsnoop

import (
	"io"
	"net/http"
	"time"
)

// Metrics holds metrics captured from CaptureMetrics.
type Metrics struct {
	// Code is the first http response code passed to the WriteHeader func of
	// the ResponseWriter. If no such call is made, a default code of 200 is
	// being used instead.
	Code int
	// Duration is the time it took to execute the handler.
	Duration time.Duration
	// Written is the number of bytes successfully written by the Write or
	// ReadFrom function of the ResponseWriter. ResponseWriters may also write
	// data to their underlaying connection directly (e.g. headers), but those
	// are not tracked. Therefor the number of Written bytes will usually match
	// the size of the response body.
	Written int64
}

// CaptureMetrics wraps the given hnd, executes it with the given w and r, and
// returns the metrics it captured from it.
func CaptureMetrics(hnd http.Handler, w http.ResponseWriter, r *http.Request) Metrics {
	// We use the updates channel and for loop below as an event loop that
	// guarantees non-overlapping execution of the closures sent through the
	// channel. This allows safe access to the m struct, even if hooks are being
	// called concurrently. The same could also be accomplished by a mutex, but
	// I've been waiting for an opportunity to try out the ideas from a few talks
	// on the topic [1][2] :).
	//
	// [1] https://www.youtube.com/watch?v=5buaPyJ0XeQ
	// [2] https://www.youtube.com/watch?v=yCbon_9yGVs

	var (
		start            = time.Now()
		m                = Metrics{Code: http.StatusOK}
		writeHeaderCount int
		updates          = make(chan func())
		done             = make(chan struct{})
		hooks            = Hooks{
			WriteHeader: func(next WriteHeaderFunc) WriteHeaderFunc {
				return func(code int) {
					// We need to do this select in every hook, otherwise we would block
					// callers from go routines that exceed the call duration of the
					// hnd.ServeHTTP call below. One may argue that this would be
					// justifiable punishment for those misbehaved callers, but I'm
					// feeling charitable today ;).
					select {
					case updates <- func() {
						if writeHeaderCount == 0 {
							m.Code = code
							writeHeaderCount++
						}
						next(code)
					}:
					case <-done:
					}
				}
			},

			Write: func(next WriteFunc) WriteFunc {
				return func(p []byte) (int, error) {
					n, err := next(p)
					select {
					case updates <- func() { m.Written += int64(n) }:
					case <-done:
					}
					return n, err
				}
			},

			ReadFrom: func(next ReadFromFunc) ReadFromFunc {
				return func(src io.Reader) (int64, error) {
					n, err := next(src)
					select {
					case updates <- func() { m.Written += n }:
					case <-done:
					}
					return n, err
				}
			},
		}
	)

	// Having to spawn an additional go routine here might be a bit unfortunate
	// from a performance perspective, but I'm not sure if it can be avoided.
	// --fg
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
