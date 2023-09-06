package gintest

import "testing"

func TestRunServer(t *testing.T) {

	result := RunServer()
	if result != nil {
		t.Errorf("rrrorrrr")
	}
}