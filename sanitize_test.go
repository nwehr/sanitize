package sanitize

import "testing"

func TestString(t *testing.T) {
	_, err := String("Nathan; drop table users")
	if err == nil {
		t.Error("expected error; got nil")
	}
}
