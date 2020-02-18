package board

import (
	"github.com/cleancode4ever/tic-tac-toe-api/internal/cell"
	"testing"
)

func TestSetEmptyBoard(t *testing.T) {
	var row, col byte = 0, 1
	var value = cell.X

	board := new(Board)
	initialBord := *board

	err := board.SetCell(row, col, value)

	if err != nil {
		t.Error(err)
	}

	if initialBord == *board {
		t.Errorf("Board state didn't change through a valid SET CELL operation: %v", board)
	}
}

func TestOutOfRangeRowAndColumn(t *testing.T) {
	var value = cell.X
	var row, col byte = Dimension, Dimension

	board := new(Board)
	initialBord := *board

	err := board.SetCell(row, col, value)

	if err == nil {
		t.Error("Setting a out-of-range cell to a valid value didn't issue any error")
	}

	if initialBord != *board {
		t.Errorf("Board state changed through an invalid operation; received: %v, expected %v", board, initialBord)
	}
}

func TestSetCellToAnInvalidValue(t *testing.T) {
	var row, col byte = 0, 1
	var value cell.Value = 0xff //Could be any value other than E (0), O (1), and X (2)

	board := new(Board)
	initialBord := *board

	err := board.SetCell(row, col, value)

	if err == nil {
		t.Error("Setting a cell to an invalid value didn't issue any error")
	}

	if initialBord != *board {
		t.Errorf("Board state changed through an invalid operation; received: %v, expected %v", board, initialBord)
	}
}

func TestSetNonEmptyBoardCell(t *testing.T) {
	var initialValue = cell.X
	var row, col byte = 0, 1
	var overwritingValue = cell.O

	board := new(Board)
	_ = board.SetCell(row, col, initialValue)
	initialBord := *board

	err := board.SetCell(row, col, overwritingValue)

	if err == nil {
		t.Errorf("<(%d, %d) = %d> was overwritten by %d", row, col, initialValue, overwritingValue)
	}

	if initialBord != *board {
		t.Errorf("Board state changed through an invalid operation; received: %v, expected %v", board, initialBord)
	}
}

func TestSetBoardToValidBoard(t *testing.T) {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.E, cell.X},
		{cell.X, cell.O, cell.O},
		{cell.O, cell.E, cell.X},
	}

	if err := new(Board).Set(&content); err != nil {
		t.Error(err)
	}
}

func TestSetBoardToInvalidBoard(t *testing.T) {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.E, cell.Value(0x22)},
		{cell.X, cell.O, cell.O},
		{cell.O, cell.E, cell.X},
	}

	if err := new(Board).Set(&content); err == nil {
		t.Error(err)
	}
}

func TestNumCellsWithValueCount(t *testing.T) {
	board := new(Board)

	_ = board.SetCell(0, 0, cell.X)
	_ = board.SetCell(1, 1, cell.X)
	_ = board.SetCell(2, 2, cell.O)
	_ = board.SetCell(1, 0, cell.O)
	_ = board.SetCell(2, 0, cell.O)

	numXCells := board.NumCellsWithValue(cell.X)
	numOCells := board.NumCellsWithValue(cell.O)

	if numXCells != 2 {
		t.Errorf("number of X cells miscalculated <%d>; expected <2>", numXCells)
	}

	if numOCells != 3 {
		t.Errorf("number of O cells miscalculated <%d>; expected <3>", numOCells)
	}
}

func TestFullRowSameValueWithAllXs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.E, cell.E},
		{cell.X, cell.X, cell.X},
		{cell.O, cell.E, cell.X},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullRowSameValue(); value != cell.X {
		t.Errorf("full row of X was not reported correctly; <%s> was given", value)
	}
}

func TestFullRowSameValueWithNoConsecutiveXOInRows(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.E, cell.E, cell.E},
		{cell.O, cell.X, cell.X},
		{cell.O, cell.E, cell.X},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullRowSameValue(); value != cell.E {
		t.Errorf("no full row of X/O yet <%s> was given; expected <%s>", value, cell.E)
	}
}

func TestFullColumnSameValueWithAllXs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.X, cell.E},
		{cell.X, cell.X, cell.X},
		{cell.O, cell.X, cell.X},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullColumnSameValue(); value != cell.X {
		t.Errorf("full column of X was not reported correctly; <%s> was given", value)
	}
}

func TestFullColumnSameValueWithNoColumnarXOs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.E, cell.X, cell.E},
		{cell.O, cell.X, cell.X},
		{cell.O, cell.E, cell.X},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullColumnSameValue(); value != cell.E {
		t.Errorf("no full column of X/O yet <%s> was given; expected <%s>", value, cell.E)
	}
}

func TestFullDiagonalSameValueWithAllXs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.X, cell.E, cell.E},
		{cell.O, cell.X, cell.O},
		{cell.O, cell.E, cell.X},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullDiagonalSameValue(); value != cell.X {
		t.Errorf("full diagonal of X was not reported correctly; <%s> was given", value)
	}
}

func TestFullDiagonalSameValueWithNoConsecutiveXOs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.X, cell.E},
		{cell.O, cell.X, cell.X},
		{cell.O, cell.E, cell.E},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullDiagonalSameValue(); value != cell.E {
		t.Errorf("no full diagonal of X/O yet <%s> was given; expected <%s>", value, cell.E)
	}
}
func TestFullAntiDiagonalSameValueWithAllXs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.X, cell.E, cell.X},
		{cell.O, cell.X, cell.O},
		{cell.X, cell.E, cell.E},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullAntiDiagonalSameValue(); value != cell.X {
		t.Errorf("full anti-diagonal of X was not reported correctly; <%s> was given", value)
	}
}

func TestFullAntiDiagonalSameValueWithNoConsecutiveXOs(t *testing.T)  {
	content := [Dimension][Dimension]cell.Value{
		{cell.O, cell.X, cell.E},
		{cell.O, cell.X, cell.X},
		{cell.O, cell.E, cell.E},
	}

	board := new(Board)
	board.Set(&content)

	if value := board.FullAntiDiagonalSameValue(); value != cell.E {
		t.Errorf("no full anti-diagonal of X/O yet <%s> was given; expected <%s>", value, cell.E)
	}
}