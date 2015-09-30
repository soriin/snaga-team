package test

import (
	"net/http"
)

type FakeResponseWriter struct {
	Calls map[string]int
	Bytes []byte
	WriteHeaderValue int
}

func NewFakeResponseWriter() *FakeResponseWriter {
	w := FakeResponseWriter{}
	w.Calls = make(map[string]int)
	w.Bytes = make([]byte, 0, 1)
	return &w
}

func (w *FakeResponseWriter) Header() http.Header {
	w.Calls["Header"] = w.Calls["Header"] + 1
	return make(http.Header)
}

func (w *FakeResponseWriter) Write(b []byte) (int, error) {
	w.Calls["Write"] = w.Calls["Write"] + 1
	w.Bytes = append(w.Bytes, b...)
	return len(b), nil
}

func (w *FakeResponseWriter) WriteHeader(a int) {
	w.Calls["WriteHeader"] = w.Calls["WriteHeader"] + 1
	w.WriteHeaderValue = a
}

func (w *FakeResponseWriter) GetOutput() []byte {
	return w.Bytes
}
