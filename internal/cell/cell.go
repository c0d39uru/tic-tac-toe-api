package cell

import "fmt"

type Value byte

const (
	E Value = iota 
	X
	O
)

func Validate(value Value) bool {
	for _, n := range Values() {
		if value == n {
			return true
		}
	}

	return false
}

func Values() []Value {
	return []Value{E, X, O}
}

func (v Value) NextTurn() Value {
	if v == X {
		return O
	}

	return X
}

func (v Value) String() string {
	switch v {
		case E:
			return fmt.Sprintf("%d (%c)", E, '␢') //0x2422
		case X: 
			return fmt.Sprintf("%d (%c)", X, 'X')
		case O:
			return fmt.Sprintf("%d (%c)", O, 'O')
		default:
			return fmt.Sprintf("%d (%s)", v, "🚫") //0x1F6AB
	}
}