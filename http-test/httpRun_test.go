package httptest

import (
	"testing"
)

func TestHttpHandle(t *testing.T) {

	if err := HttpHandle(); err != nil {
		t.Errorf("error")
	}
}