package auth

import (
	"errors"
	"testing"
)

func TestNewAuthError(t *testing.T) {
	e := NewAuthError(100, errors.New("some error"))
	if "some error" != e.Error() {
		t.Fatalf("Error method returned: %s", e.Error())
	}
	if 100 != e.ErrorCode() {
		t.Fatalf("ErrorCode method returned %d", e.ErrorCode())
	}
}
