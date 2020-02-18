package state

import "fmt"

type State struct {
	Code    byte   `json:"code"`
	Message string `json:"message"`
}

const (
	NotStarted byte = iota
	InProgress
	GameOverXWon
	GameOverOWon
	GameOverDraw
	InvalidBoardState
)

func New(code byte, message string) (*State, error) {
	if !validate(code) {
		return nil, fmt.Errorf("invalid state given")
	}

	if "" == message {
		message = getDefaultMessage(code, message)
	}

	return &State{
		Code:    code,
		Message: message,
	}, nil
}

func getDefaultMessage(code byte, message string) string {
	switch code {
	case NotStarted:
		message = "not started"
	case InProgress:
		message = "in progress"
	case GameOverXWon:
		message = "game over: x won"
	case GameOverOWon:
		message = "game over: o won"
	case GameOverDraw:
		message = "game over: draw"
	case InvalidBoardState:
		message = "invalid board state"
	}

	return message
}

func values() []byte {
	return []byte{NotStarted, InProgress, GameOverXWon, GameOverOWon, GameOverDraw, InvalidBoardState}
}

func validate(code byte) bool {
	for _, n := range values() {
		if code == n {
			return true
		}
	}

	return false
}
