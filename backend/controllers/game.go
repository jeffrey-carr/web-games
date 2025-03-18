package controllers

import (
	"binoku/entities"
	"binoku/utils"
	"fmt"
	"slices"
)

// GameManager represents a game manager
type GameManager interface {
	GenerateBoard(int) (entities.Game, error)
	ValidateBoard(board [][]int) (bool, entities.InvalidBoardHint)
}

type gameManager struct{}

// NewGameManager is the constructor for a GameManager
func NewGameManager() GameManager {
	return &gameManager{}
}

func (gm gameManager) GenerateBoard(size int) (entities.Game, error) {
	board := generateGameBoard(size)

	return entities.Game{Board: board}, nil
}

func (gm gameManager) ValidateBoard(board [][]int) (bool, entities.InvalidBoardHint) {
	// Confirm there are no empty spaces
	for i, row := range board {
		if slices.Contains(row, -1) {
			return false, entities.InvalidBoardHint{
				Rows: []int{i},
			}
		}
	}

	return boardIsValid(board)
}

// Generate a fully valid board
func generateGameBoard(n int) [][]int {
	var board [][]int
	// Fill board with -1
	for range n {
		row := make([]int, n)
		for i := range row {
			row[i] = -1
		}
		board = append(board, row)
	}

	backtrackFill(board, 0, 0)

	fmt.Println("Generated board:")
	for _, row := range board {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}

		fmt.Printf("(%d)", utils.SumSlice(row))
		fmt.Println()
	}

	return backtrackSolve(board)
}

func backtrackFill(board [][]int, row int, col int) bool {
	n := len(board)
	if row == n {
		return true
	}

	// Randomly pick a 0 or 1
	options := []int{0, 1}
	options = utils.ShuffleSlice(options)

	nextCol := col + 1
	nextRow := row
	if nextCol >= n {
		nextCol = 0
		nextRow++
	}

	for _, possibleValue := range options {
		if valueIsValid(board, row, col, possibleValue) {
			board[row][col] = possibleValue
			if backtrackFill(board, nextRow, nextCol) {
				return true
			}
			// It didn't work, undo move
			board[row][col] = -1
		}
	}

	return false
}

// boardIsValid validates that a board is correct. Returns the
func boardIsValid(board [][]int) (bool, entities.InvalidBoardHint) {
	// The game has 3 rules:
	// 1. There must be an equal number of 1's and 0's in each row/column
	// 2. There cannot be more than 2 consecutive values next to each other in each row/column
	// 3. There cannot be any identical rows or any identical columns
	for rowIndex, row := range board {
		isRuleOneValid := validateRuleOne(row)
		if !isRuleOneValid {
			return false, entities.InvalidBoardHint{
				Rows: []int{rowIndex},
			}
		}
		isRuleTwoValid := validateRuleTwo(row)
		if !isRuleTwoValid {
			return false, entities.InvalidBoardHint{
				Rows: []int{rowIndex},
			}
		}
		isRuleThreeValid, offendingRowIndex := validateRuleThree(rowIndex, board)
		if !isRuleThreeValid {
			return false, entities.InvalidBoardHint{
				Rows: []int{rowIndex, offendingRowIndex},
			}
		}
	}

	// Transpose matrix 90 degrees and check the columns
	transposed := utils.TransposeMatrix(board)
	for colIndex, col := range transposed {
		isRuleOneValid := validateRuleOne(col)
		if !isRuleOneValid {
			return false, entities.InvalidBoardHint{
				Cols: []int{colIndex},
			}
		}
		isRuleTwoValid := validateRuleTwo(col)
		if !isRuleTwoValid {
			return false, entities.InvalidBoardHint{
				Cols: []int{colIndex},
			}
		}
		isRuleThreeValid, offendingColIndex := validateRuleThree(colIndex, transposed)
		if !isRuleThreeValid {
			return false, entities.InvalidBoardHint{
				Rows: []int{colIndex, offendingColIndex},
			}
		}
	}

	return true, entities.InvalidBoardHint{}

	// for n := range board {
	// 	rIsValid := rowIsValid(utils.DuplicateBoard(board), n)
	// 	cIsValid := colIsValid(utils.DuplicateBoard(board), n)

	// 	if !rIsValid || !cIsValid {
	// 		return false
	// 	}
	// }
	// return true
}

// 1. There must be an equal number of 1's and 0's in each row/column
func validateRuleOne(row []int) bool {
	size := len(row)
	rowZeros, rowOnes := 0, 0
	for _, value := range row {
		if value == 0 {
			rowZeros++
		}
		if value == 1 {
			rowOnes++
		}
	}

	// Rule 1
	if rowZeros > size/2 || rowOnes > size/2 {
		return false
	}

	return true
}

// 2. There cannot be more than 2 consecutive values next to each other in each row/column
func validateRuleTwo(row []int) bool {
	// Rule 2
	for i, value := range row {
		if i < 2 {
			continue
		}

		if value < 0 {
			continue
		}

		if value == row[i-1] && value == row[i-2] {
			return false
		}
	}

	return true
}

// 3. There cannot be any identical rows or any identical columns
// Returns whether the rule is valid. If not, provides the index of
// the first matched column
func validateRuleThree(rowIndex int, board [][]int) (bool, int) {
	if rowIndex < 0 || rowIndex > len(board) {
		return false, -1
	}

	row := board[rowIndex]
	for innerRowIndex, innerRow := range board {
		if innerRowIndex == rowIndex {
			continue
		}

		if slices.EqualFunc(row, innerRow, func(item int, innerItem int) bool {
			if item < 0 || innerItem < 0 {
				return false
			}

			return item == innerItem
		}) {
			return false, innerRowIndex
		}
	}

	return true, -1
}

func valueIsValid(toValidate [][]int, row int, col int, value int) bool {
	board := utils.DuplicateBoard(toValidate)
	board[row][col] = value

	isValid, _ := boardIsValid(board)
	return isValid
}

func backtrackSolve(board [][]int) [][]int {
	n := len(board)
	// To convert a board into a puzzle:
	// Step 1: Remove a number
	// Step 2: Check if board is still uniquely solvable

	// To minimize going down bad routes and for a better user experience,
	// we will attempt to take an equal amount from each quadrant, so we
	// will visit the quadrants in a round-robin fashion
	var topLeft, topRight, bottomLeft, bottomRight []entities.Coordinate
	for row := range n {
		for col := range n {
			coord := entities.Coordinate{Col: col, Row: row}

			top := row < n/2
			left := col < n/2
			if top {
				if left {
					topLeft = append(topLeft, coord)
				} else {
					topRight = append(topRight, coord)
				}
			} else {
				if left {
					bottomLeft = append(bottomLeft, coord)
				} else {
					bottomRight = append(bottomRight, coord)
				}
			}
		}
	}
	// Shuffle each to make sure we are visiting cells randomly
	topLeft = utils.ShuffleSlice(topLeft)
	bottomLeft = utils.ShuffleSlice(bottomLeft)
	topRight = utils.ShuffleSlice(topRight)
	bottomRight = utils.ShuffleSlice(bottomRight)
	// Now merge them into a new copy
	coords := []entities.Coordinate{}
	for i := range n * n {
		var coord entities.Coordinate
		if i%4 == 0 && len(topLeft) > 0 {
			coord = topLeft[0]
			topLeft = topLeft[1:]
		} else if i%4 == 1 && len(topRight) > 0 {
			coord = topRight[0]
			topRight = topRight[1:]
		} else if i%4 == 2 && len(bottomLeft) > 0 {
			coord = bottomLeft[0]
			bottomLeft = bottomLeft[1:]
		} else {
			coord = bottomRight[0]
			bottomRight = bottomRight[1:]
		}

		coords = append(coords, coord)
	}

	emptySpaces := []entities.Coordinate{}

	// We can start by removing the first space, since there is nothing to solve
	coord := coords[0]
	coords = coords[1:]
	board[coord.Row][coord.Col] = -1
	emptySpaces = append(emptySpaces, coord)

	batchSize := 3
	for len(coords) > 6 {
		// Step 1
		removedValues, removedCoords := removeCoords(board, coords, batchSize)
		emptySpaces = append(emptySpaces, removedCoords...)

		coords = coords[batchSize:]

		if !isUniquelySolvable(board, emptySpaces) {
			// If it is not uniquely solvable, try putting back one number at a time
			foundUnique := false
			for i := range batchSize {
				coord := removedCoords[i]
				value := removedValues[i]

				board[coord.Row][coord.Col] = value
				emptySpaces = slices.DeleteFunc(emptySpaces, func(emptyCoord entities.Coordinate) bool {
					return emptyCoord.Col == coord.Col && emptyCoord.Row == coord.Row
				})

				if isUniquelySolvable(board, emptySpaces) {
					foundUnique = true
					break
				}
			}

			if !foundUnique {
				return board
			}
		}
	}

	return board
}

func removeCoords(board [][]int, coords []entities.Coordinate, n int) ([]int, []entities.Coordinate) {
	if n > len(coords) {
		return nil, nil
	}

	removedValues := []int{}
	removedCoords := []entities.Coordinate{}

	for i := range n {
		coord := coords[i]
		removedCoords = append(removedCoords, coord)

		value := board[coord.Row][coord.Col]
		board[coord.Row][coord.Col] = -1
		removedValues = append(removedValues, value)
	}

	return removedValues, removedCoords
}

func isUniquelySolvable(board [][]int, emptySpaces []entities.Coordinate) bool {
	count := 0
	countSolutions(board, emptySpaces, &count)
	return count == 1
}

// Solves a board and returns if it was successful
func countSolutions(board [][]int, emptySpaces []entities.Coordinate, count *int) {
	// Base case
	if len(emptySpaces) == 0 {
		*count++
		return
	}

	if *count > 1 {
		return
	}

	coord, emptySpaces := emptySpaces[0], emptySpaces[1:]

	if valueIsValid(board, coord.Row, coord.Col, 0) {
		bCpy := utils.DuplicateBoard(board)
		bCpy[coord.Row][coord.Col] = 0
		countSolutions(bCpy, emptySpaces, count)
	}
	if valueIsValid(board, coord.Row, coord.Col, 1) {
		bCpy := utils.DuplicateBoard(board)
		bCpy[coord.Row][coord.Col] = 1
		countSolutions(bCpy, emptySpaces, count)
	}

	return
}
