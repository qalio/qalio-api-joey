package joey

import "testing"

func TestInit(t *testing.T) {
	expected := "joey module initialized"
	if got := Init(); got != expected {
		t.Errorf("Init() = %q, want %q", got, expected)
	}
}
