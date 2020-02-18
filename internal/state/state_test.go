package state

import (
	"testing"
)

func TestInstantiatingWithValidValues(t *testing.T) {
	for _, value := range values() {
		_, err := New(value, "")

		if err != nil {
			t.Errorf("valid state <%d> did not result in a valid State", value)
		}
	}
}

func TestInstantiatingWithInvalidValues(t *testing.T) {
	const maxByte = ^byte(0)

	for i := byte(len(values())); i < maxByte; i++ {
		_, err := New(i, "")

		if err == nil {
			t.Errorf("invalid state <%d> did not result in an error", i)
		}
	}
}

func TestInstantiatingWithCustomMessage(t *testing.T) {
	code := InvalidBoardState
	message := "some custom message"
	state, err := New(code, message)

	if err != nil || state.Code != code || state.Message != message {
		t.Errorf("valid state <%d> and message <%s> did not result in a valid State; <%v> given", code, message, state)
	}
}