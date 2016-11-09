package httpsnoop

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// @TODO hijacker, ReaderFrom ?

func TestSnoopResponserWriter_interfaces(t *testing.T) {
	tests := []struct {
		W                 http.ResponseWriter
		WantFlusher       bool
		WantCloseNotifier bool
	}{
		{
			W:                 struct{ http.ResponseWriter }{},
			WantFlusher:       false,
			WantCloseNotifier: false,
		},
		{
			W: struct {
				http.ResponseWriter
				http.Flusher
			}{},
			WantFlusher:       true,
			WantCloseNotifier: false,
		},
		{
			W: struct {
				http.ResponseWriter
				http.CloseNotifier
			}{},
			WantFlusher:       false,
			WantCloseNotifier: true,
		},
		{
			W: struct {
				http.ResponseWriter
				http.Flusher
				http.CloseNotifier
			}{},
			WantFlusher:       true,
			WantCloseNotifier: true,
		},
	}
	for _, test := range tests {
		sw := SnoopResponseWriter(test.W, Hooks{})
		if _, got := sw.(http.CloseNotifier); got != test.WantCloseNotifier {
			t.Errorf("got=%t want=%t", got, test.WantCloseNotifier)
		}
		if _, got := sw.(http.Flusher); got != test.WantFlusher {
			t.Errorf("got=%t want=%t", got, test.WantFlusher)
		}
	}
}

func TestSnoopResponserWriter_integration(t *testing.T) {
	tests := []struct {
		Name     string
		Handler  http.Handler
		Hooks    Hooks
		WantCode int
		WantBody []byte
	}{
		{
			Name: "WriteHeader (no hook)",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}),
			WantCode: http.StatusNotFound,
		},
		{
			Name: "WriteHeader",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}),
			Hooks: Hooks{
				WriteHeader: func(next WriteHeaderFunc) WriteHeaderFunc {
					return func(code int) {
						if code != http.StatusNotFound {
							t.Errorf("got=%d want=%d", code, http.StatusNotFound)
						}
						next(http.StatusForbidden)
					}
				},
			},
			WantCode: http.StatusForbidden,
		},

		{
			Name: "Write (no hook)",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("foo"))
			}),
			WantCode: http.StatusOK,
			WantBody: []byte("foo"),
		},
		{
			Name: "Write (rewrite hook)",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if n, err := w.Write([]byte("foo")); err != nil {
					t.Errorf("got=%s", err)
				} else if got, want := n, len("foobar"); got != want {
					t.Errorf("got=%d want=%d", got, want)
				}
			}),
			Hooks: Hooks{
				Write: func(next WriteFunc) WriteFunc {
					return func(p []byte) (int, error) {
						if string(p) != "foo" {
							t.Errorf("%s", p)
						}
						return next([]byte("foobar"))
					}
				},
			},
			WantCode: http.StatusOK,
			WantBody: []byte("foobar"),
		},
		{
			Name: "Write (error hook)",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if n, err := w.Write([]byte("foo")); n != 0 {
					t.Errorf("got=%d want=%d", n, 0)
				} else if err != io.EOF {
					t.Errorf("got=%s want=%s", err, io.EOF)
				}
			}),
			Hooks: Hooks{
				Write: func(next WriteFunc) WriteFunc {
					return func(p []byte) (int, error) {
						if string(p) != "foo" {
							t.Errorf("%s", p)
						}
						return 0, io.EOF
					}
				},
			},
			WantCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		func() {
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				sw := SnoopResponseWriter(w, test.Hooks)
				test.Handler.ServeHTTP(sw, r)
			})
			s := httptest.NewServer(h)
			defer s.Close()
			res, err := http.Get(s.URL)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			gotBody, err := ioutil.ReadAll(res.Body)
			if res.StatusCode != test.WantCode {
				t.Errorf("got=%d want=%d", res.StatusCode, test.WantCode)
			} else if !bytes.Equal(gotBody, test.WantBody) {
				t.Errorf("got=%s want=%s", gotBody, test.WantBody)
			}
		}()
	}
}

func TestSnoopHandlerMetrics(t *testing.T) {
	tests := []struct {
		Handler      http.Handler
		WantDuration time.Duration
		WantWritten  int64
		WantCode     int
	}{
		{
			Handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			WantCode: http.StatusOK,
		},
		{
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
	}

	for _, test := range tests {
		func() {
			ch := make(chan Metrics, 1)
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ch <- SnoopMetrics(test.Handler, w, r)
			})
			s := httptest.NewServer(h)
			defer s.Close()
			res, err := http.Get(s.URL)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			m := <-ch
			if m.Code != test.WantCode {
				t.Errorf("got=%d want=%d", m.Code, test.WantCode)
			} else if m.Duration < test.WantDuration {
				t.Errorf("got=%s want=%s", m.Duration, test.WantDuration)
			} else if m.Written < test.WantWritten {
				t.Errorf("got=%d want=%d", m.Written, test.WantWritten)
			}
		}()
	}
}
