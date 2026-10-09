package httpsnoop

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"testing/iotest"
	"time"
)

func TestCaptureMetrics(t *testing.T) {
	// Some of the edge cases tested below cause the net/http pkg to log some
	// messages that add a lot of noise to the `go test -v` output, so we discard
	// the log here.
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	tests := []struct {
		Name         string
		Handler      http.Handler
		WantDuration time.Duration
		WantWritten  int64
		WantCode     int
		WantClient   int
		WantErr      string
	}{
		{
			Name:     "simple",
			Handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			WantCode: http.StatusOK,
		},
		{
			Name: "headers and body",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("foo"))
				w.Write([]byte("bar"))
				time.Sleep(25 * time.Millisecond)
			}),
			WantCode:     http.StatusBadRequest,
			WantWritten:  6,
			WantDuration: 25 * time.Millisecond,
		},
		{
			Name: "header after body",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("foo"))
				w.WriteHeader(http.StatusNotFound)
			}),
			WantCode: http.StatusOK,
		},
		{
			Name: "reader",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				rrf := w.(io.ReaderFrom)
				rrf.ReadFrom(strings.NewReader("reader from is ok"))
			}),
			WantWritten: 17,
			WantCode:    http.StatusOK,
		},
		{
			Name: "readfrom error then 502",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := io.Copy(w, iotest.ErrReader(errors.New("upstream failed")))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadGateway)
				}
			}),
			WantCode:   http.StatusBadGateway,
			WantClient: http.StatusBadGateway,
		},
		{
			Name: "readfrom empty then 204",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// A reader without WriteTo forces io.Copy through ReadFrom.
				_, _ = io.Copy(w, eofReader{})
				w.WriteHeader(http.StatusNoContent)
			}),
			WantCode:   http.StatusNoContent,
			WantClient: http.StatusNoContent,
		},
		{
			Name: "string writer",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = io.WriteString(w, "write string")
			}),
			WantWritten: int64(len("write string")),
			WantCode:    http.StatusOK,
		},
		{
			Name: "empty panic",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("oh no")
			}),
			WantCode:     http.StatusOK, // code is not written so is our default
			WantDuration: 1,             // confirm non-zero
			WantErr:      "EOF",
		},
		{
			Name: "panic after partial response",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("failed to execute"))
				panic("oh no")
			}),
			WantCode:     http.StatusInternalServerError,
			WantDuration: 1, // confirm non-zero
			WantWritten:  17,
			WantErr:      "EOF",
		},
	}

	for i, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ch := make(chan Metrics, 1)
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				m := Metrics{Code: http.StatusOK}
				defer func() {
					ch <- m
				}()
				m.CaptureMetrics(w, func(ww http.ResponseWriter) {
					test.Handler.ServeHTTP(ww, r)
				})
			})
			s := httptest.NewServer(h)
			defer s.Close()
			res, err := http.Get(s.URL)
			if !errContains(err, test.WantErr) {
				t.Errorf("test %d: got=%s want=%s", i, err, test.WantErr)
			}
			if err == nil {
				defer res.Body.Close()
				if test.WantClient != 0 && res.StatusCode != test.WantClient {
					t.Errorf("test %d: client got=%d want=%d", i, res.StatusCode, test.WantClient)
				}
			}
			m := <-ch
			if m.Code != test.WantCode {
				t.Errorf("test %d: got=%d want=%d", i, m.Code, test.WantCode)
			} else if m.Duration < test.WantDuration {
				t.Errorf("test %d: got=%s want=%s", i, m.Duration, test.WantDuration)
			} else if m.Written < test.WantWritten {
				t.Errorf("test %d: got=%d want=%d", i, m.Written, test.WantWritten)
			}
		})
	}
}

type eofReader struct{}

func (eofReader) Read([]byte) (int, error) { return 0, io.EOF }

func errContains(err error, s string) bool {
	var errS string
	if err == nil {
		errS = ""
	} else {
		errS = err.Error()
	}
	return strings.Contains(errS, s)
}
