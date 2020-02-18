package game

import (
	"errors"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/board"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/cell"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/state"
	"math"
)

type Status struct {
	Turn  cell.Value   `json:"turn"`
	State *state.State `json:"state"`
}
type Game struct {
	Status *Status      `json:"status"`
	Board  *board.Board `json:"board"`
}

func New() *Game {
	initState, _ := state.New(state.NotStarted, "")

	return &Game{
		Status: &Status{
			Turn:  cell.X,
			State: initState,
		},
		Board: new(board.Board),
	}
}

func (g *Game) SetBoard(content *[board.Dimension][board.Dimension]cell.Value, row, col byte) error {
	err := g.initBoard(content)

	if err != nil {
		return err
	}

	if g.Status.Turn == cell.E {
		return errors.New(g.Status.State.Message)
	}

	err = g.Board.SetCell(row, col, g.Status.Turn)

	if err != nil {
		g.updateStatusWithError(err)
	} else {
		g.updateStatus()
	}

	return err
}

func (g *Game) initBoard(content *[board.Dimension][board.Dimension]cell.Value) error {
	gameBoard := new(board.Board)

	if err := gameBoard.Set(content); err != nil {
		g.updateStatusWithError(err)
		return err
	}

	g.Board = gameBoard

	g.updateStatus()

	return nil
}

func (g *Game) updateStatusWithError(e error) {
	g.Status.Turn = cell.E
	g.Status.State, _ = state.New(state.InvalidBoardState, e.Error())
}

func (g *Game) updateStatus() {
	var nextState *state.State
	numXs := g.Board.NumCellsWithValue(cell.X)
	numOs := g.Board.NumCellsWithValue(cell.O)

	switch {
	case byte(math.Abs(float64(numXs) - float64(numOs))) >= 2:
		g.Status.Turn = cell.E
		nextState, _ = state.New(state.InvalidBoardState, "")

	case g.isDraw():
		g.Status.Turn = cell.E
		nextState, _ = state.New(state.GameOverDraw, "")

	case g.xWon():
		g.Status.Turn = cell.E
		nextState, _ = state.New(state.GameOverXWon, "")

	case g.oWon():
		g.Status.Turn = cell.E
		nextState, _ = state.New(state.GameOverOWon, "")

	default:
		nextState, _ = state.New(state.InProgress, "")
		if numXs > numOs {
			g.Status.Turn = cell.O
		} else {
			g.Status.Turn = cell.X
		}
	}

	g.Status.State = nextState
}

func (g *Game) xWon() bool {
	return g.isWinner(cell.X)
}

func (g *Game) oWon() bool {
	return g.isWinner(cell.O)
}

func (g *Game) isDraw() bool {
	return 0 == g.Board.NumCellsWithValue(cell.E) && !g.xWon() && !g.oWon()
}

func (g *Game) isWinner(value cell.Value) bool {
	return value == g.Board.FullRowSameValue() ||
		value == g.Board.FullColumnSameValue() ||
		value == g.Board.FullDiagonalSameValue() ||
		value == g.Board.FullAntiDiagonalSameValue()
}
