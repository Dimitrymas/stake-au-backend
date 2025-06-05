package customerrors

import (
	"errors"
	"testing"
)

func TestPartialAccountsError(t *testing.T) {
	err := NewPartialAccountsError(2, 1)
	var pe *partialAccountsError
	if !errors.As(err, &pe) {
		t.Fatalf("unexpected type")
	}
	if pe.Created != 2 || pe.NotCreated != 1 {
		t.Errorf("wrong values: %#v", pe)
	}
	if !errors.Is(err, ErrCreatePartialAccounts) {
		t.Error("errors.Is failed")
	}
	if pe.Error() == "" {
		t.Error("empty error message")
	}
}
