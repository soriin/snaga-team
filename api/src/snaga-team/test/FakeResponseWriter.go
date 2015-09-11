package test

import (
	"net/http"
)

type FakeResponseWriter struct {
	Calls map[string]int
}

func NewFakeResponseWriter() *FakeResponseWriter {
	w := FakeResponseWriter{}
	w.Calls = make(map[string]int)
	return &w
}

func (w *FakeResponseWriter) Header() http.Header {
	w.Calls["Header"] = w.Calls["Header"] + 1
	return make(http.Header)
}

func (w *FakeResponseWriter) Write(b []byte) (int, error) {
	w.Calls["Write"] = w.Calls["Write"] + 1
	return 0, nil
}

func (w *FakeResponseWriter) WriteHeader(int) {
	w.Calls["WriteHeader"] = w.Calls["WriteHeader"] + 1
}
