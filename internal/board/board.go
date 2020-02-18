package board

import (
	"fmt"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/cell"
)

const Dimension = 3

type Board struct {
	Content [Dimension][Dimension]cell.Value `json:"content"`
}

func (b *Board) SetCell(row, col byte, value cell.Value) error {
	currentCellValue, err := b.cellValueAt(row, col)

	if err != nil {
		return err
	}

	if !cell.Validate(value) {
		return fmt.Errorf("invalid cell value <%d> - valid values are: %v", value, cell.Values())
	}

	if currentCellValue != cell.E {
		return fmt.Errorf("non-empty cell (%d, %d) = %d cannot be set to %d", row, col, currentCellValue, value)
	}

	b.Content[row][col] = value

	return nil
}

func (b *Board) Set(content *[Dimension][Dimension]cell.Value) error {
	var i, j byte

	for i = 0; i < Dimension; i++ {
		for j = 0; j < Dimension; j++ {
			err := b.SetCell(i, j, content[i][j])

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Board) NumCellsWithValue(value cell.Value) byte {
	var count byte = 0

	for i := 0; i < Dimension; i++ {
		for j := 0; j < Dimension; j++ {
			if b.Content[i][j] == value {
				count++
			}
		}
	}

	return count
}

func (b *Board) cellValueAt(row, col byte) (cell.Value, error) {
	if !isDimensionValid(row) || !isDimensionValid(col) {
		return cell.E, fmt.Errorf("out-of-bound coordination given (%d, %d)", row, col)
	}

	return b.Content[row][col], nil
}

func (b *Board) FullRowSameValue() cell.Value {
	var i byte

	for i = 0; i < Dimension; i++ {
		leftVal, _ := b.cellValueAt(i, byte(0))
		middleVal, _ := b.cellValueAt(i, byte(1))
		rightVal, _ := b.cellValueAt(i, byte(2))

		if leftVal == cell.E || middleVal == cell.E || rightVal == cell.E {
			continue
		}

		if leftVal == middleVal && middleVal == rightVal {
			return leftVal
		}
	}

	return cell.E
}

func (b *Board) FullColumnSameValue() cell.Value {
	var i byte

	for i = 0; i < Dimension; i++ {
		topVal, _ := b.cellValueAt(byte(0), i)
		middleVal, _ := b.cellValueAt(byte(1), i)
		bottomVal, _ := b.cellValueAt(byte(2), i)

		if topVal == cell.E || middleVal == cell.E || bottomVal == cell.E {
			continue
		}

		if topVal == middleVal && middleVal == bottomVal {
			return topVal
		}
	}

	return cell.E
}

func (b *Board) FullDiagonalSameValue() cell.Value {
	topVal, _ := b.cellValueAt(byte(0), byte(0))
	middleVal, _ := b.cellValueAt(byte(1), byte(1))
	bottomVal, _ := b.cellValueAt(byte(2), byte(2))

	if topVal == middleVal && middleVal == bottomVal {
		return topVal
	}

	return cell.E
}

func (b *Board) FullAntiDiagonalSameValue() cell.Value {
	topVal, _ := b.cellValueAt(byte(0), byte(2))
	middleVal, _ := b.cellValueAt(byte(1), byte(1))
	bottomVal, _ := b.cellValueAt(byte(2), byte(0))

	if topVal == middleVal && middleVal == bottomVal {
		return topVal
	}

	return cell.E
}

func isDimensionValid(dimension byte) bool {
	return dimension < Dimension
}
