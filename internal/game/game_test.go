package game

import (
	"github.com/cleancode4ever/tic-tac-toe-api/internal/board"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/cell"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/state"
	"testing"
)

func TestGameStartsWithSaneInitialValues(t *testing.T) {
	game := New()
	initialState, _ := state.New(state.NotStarted, "")

	if *game.Status.State != *initialState {
		t.Errorf("bad game initial State <%v>; expected <%v>", game.Status.State, *initialState)
	}

	if initialTurn := cell.X; game.Status.Turn != initialTurn {
		t.Errorf("bad game initial Turn <%s>; expected <%s>", game.Status.Turn, initialTurn)
	}

	if emptyBoard := new(board.Board); *game.Board != *emptyBoard {
		t.Errorf("bad game initial Board <%v>; expected <%v>", *game.Board, *emptyBoard)
	}
}

func TestSetBoardWithInvalidBoard(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.E, cell.Value(0x50), cell.E},
		{cell.E, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
	}

	assertSetBoard(&boardArray, cell.E, state.InvalidBoardState, 0x0, 0x1, t)
}

func TestSetBoardWithFullBoard(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.X, cell.O},
		{cell.O, cell.O, cell.X},
		{cell.X, cell.O, cell.X},
	}

	assertSetBoard(&boardArray, cell.E, state.GameOverDraw, 0x0, 0x1, t)
}

func TestSetBoardWithOccupiedCell(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.E, cell.X, cell.O},
		{cell.O, cell.E, cell.X},
		{cell.X, cell.O, cell.X},
	}

	assertSetBoard(&boardArray, cell.E, state.InvalidBoardState, 0x0, 0x1, t)
}

func TestSetBoardWithLegalMove(t *testing.T) {
	game := New()
	var row, col byte = 0x0, 0x1
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.E, cell.E, cell.O},
		{cell.O, cell.E, cell.X},
		{cell.X, cell.O, cell.X},
	}

	err := game.SetBoard(&boardArray, row, col)

	if err != nil {
		t.Errorf("setting game Board to <%v> errored out", boardArray)
	}

	assertStatus(*game.Status, cell.O, state.InProgress, "", t)
}

func TestSetBoardWithValidBoard(t *testing.T) {
	var row, col byte = 0x0, 0x1
	game := New()
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.E, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
	}

	err := game.SetBoard(&boardArray, row, col)

	if err != nil {
		t.Errorf("setting (%d, %d) of an empty game <%v> to %v errored out", row, col, boardArray, cell.X)
	}

	assertStatus(*game.Status, cell.O, state.InProgress, "", t)
}

func TestSetGameBoardToAnArrayWithOneX(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
	}

	assertInitBoard(&boardArray, cell.O, state.InProgress, "", t)
}

func TestSetGameBoardToAnArrayWithOneXAndOneO(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.E, cell.E},
		{cell.E, cell.E, cell.E},
		{cell.E, cell.E, cell.O},
	}

	assertInitBoard(&boardArray, cell.X, state.InProgress, "", t)
}

func TestSetGameBoardToAnArrayWithFourXsAndOneO(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.X, cell.O},
		{cell.E, cell.E, cell.E},
		{cell.E, cell.X, cell.X},
	}

	assertInitBoard(&boardArray, cell.E, state.InvalidBoardState, "", t)
}

func TestSetGameBoardToAnArrayWhenNoEmptySpace(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.X, cell.O},
		{cell.O, cell.O, cell.X},
		{cell.X, cell.X, cell.O},
	}

	assertInitBoard(&boardArray, cell.E, state.GameOverDraw, "", t)
}

func TestIsWinnerWith3ConsecutiveXs(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.X, cell.O},
		{cell.O, cell.X, cell.E},
		{cell.E, cell.O, cell.X},
	}

	assertInitBoard(&boardArray, cell.E, state.GameOverXWon, "", t)
}

func TestIsWinnerWith3ConsecutiveOs(t *testing.T) {
	boardArray := [board.Dimension][board.Dimension]cell.Value{
		{cell.X, cell.X, cell.O},
		{cell.E, cell.X, cell.O},
		{cell.E, cell.O, cell.O},
	}

	assertInitBoard(&boardArray, cell.E, state.GameOverOWon, "", t)
}

func assertInitBoard(boardArray *[board.Dimension][board.Dimension]cell.Value, expectedNextTurn cell.Value, expectedStateCode byte, expectedStateMessage string, t *testing.T) {
	game := New()

	if err := game.initBoard(boardArray); err != nil {
		t.Errorf("setting game Board to <%v> errored out <%v>", boardArray, err)
	}

	assertStatus(*game.Status, expectedNextTurn, expectedStateCode, expectedStateMessage, t)
}

func assertSetBoard(boardArray *[board.Dimension][board.Dimension]cell.Value, expectedNextTurn cell.Value, expectedStateCode, row, col byte, t *testing.T) {
	game := New()
	err := game.SetBoard(boardArray, row, col)

	if err == nil {
		t.Errorf("setting game Board to <%v> did not error out", boardArray)
	}

	assertStatus(*game.Status, expectedNextTurn, expectedStateCode, err.Error(), t)
}

func assertStatus(status Status, expectedTurn cell.Value, expectedStateCode byte, expectedStateMessage string, t *testing.T) {
	expectedState, err := state.New(expectedStateCode, expectedStateMessage)

	if err != nil {
		t.Errorf("bad state code given: %d", expectedStateCode)
	}

	if status.Turn != expectedTurn {
		t.Errorf("bad turn calculation; received: <%v> expected: <%v>", status.Turn, expectedTurn)
	}

	if *status.State != *expectedState {
		t.Errorf("bad State calculation; received: <%v> expected: <%v>", status.State, expectedState)
	}
}
