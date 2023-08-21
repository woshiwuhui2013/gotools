package runtime_tool

import "testing"

func TestPath(t *testing.T) {
	name := "aaaa"
	want := "testa/aaaa"
	result := executablePath(name)
	if result != want {
		t.Errorf("not ok name %s", name)
	}

}
