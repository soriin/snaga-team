package helpers

import (
	"testing"

	"snaga-team/test"
)

func TestSendError(t *testing.T) {
	w := test.NewFakeResponseWriter()

	SendError(w, "errorText", 1)

	if w.Calls["Write"] == 0 {
		t.Errorf("Expected Calls to FakeResponseWriter.Write to be greater than 0 but was %v", w.Calls["Write"])
	}

	if w.WriteHeaderValue != 1 {
		t.Errorf("Expected WriteHeaderValue to be than 1 but was %v", w.WriteHeaderValue)
	}
}
