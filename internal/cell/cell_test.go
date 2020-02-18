package cell

import (
	"fmt"
	"testing"
)

func TestValidateWithInvalidValues(t *testing.T) {
	const maxByte = ^Value(0)

	for i := Value(len(Values())); i < maxByte; i++ {
		if Validate(i) != false {
			t.Errorf("invalid cell value <%d> was validated", i)
		}
	}
}

func TestValidateWithValidValues(t *testing.T) {
	for _, value := range Values() {
		if Validate(value) != true {
			t.Errorf("valid cell value <%d> was invalidated", value)
		}
	}
}

func TestNextTurn(t *testing.T) {
	var data = [...]struct {
		in  Value
		out Value
	}{
		{O, X},
		{X, O},
		{E, X},
	}

	for i := 0; i < len(data); i++ {
		v := Value(data[i].in).NextTurn()

		if v != data[i].out {
			t.Errorf("wrong calculation of the NextTurn for %v; (%v) given, (%v) expected", data[i].in, v, data[i].out)
		}
	}
}

func TestString(t *testing.T) {
	if s := fmt.Sprintf("%s", X); s != "1 (X)" {
		t.Errorf("bad stringification of X (%s)", s)
	}

	if s := fmt.Sprintf("%s", O); s != "2 (O)" {
		t.Errorf("bad stringification of O (%s)", s)
	}

	if s := fmt.Sprintf("%s", E); s != "0 (␢)" {
		t.Errorf("bad stringification of E (%s)", s)
	}

	if s := fmt.Sprintf("%s", Value(0xA1)); s != "161 (🚫)" {
		t.Errorf("bad stringification of invalid characters (%s)", s)
	}
}