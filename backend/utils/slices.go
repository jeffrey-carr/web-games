package utils

import (
	"math/rand"
	"slices"
)

// Any returns true if any elements in slice return true in the
// provided function
func Any[T any](slice []T, f func(x T) bool) bool {
	for _, x := range slice {
		if f(x) {
			return true
		}
	}

	return false
}

// ShuffleSlice takes in a slice and returns the
// slice shuffled
func ShuffleSlice[T any](slice []T) []T {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}

// SumSlice sums the values in a slice of integers
func SumSlice(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}

	return total
}

// DuplicateBoard creates a deep copy of the provided game board
func DuplicateBoard(board [][]int) [][]int {
	var dupe [][]int
	for _, row := range board {
		dupeRow := make([]int, len(row))
		copy(dupeRow, row)
		dupe = append(dupe, dupeRow)
	}

	return dupe
}

// IsEqualSlices confirms if two slices have the same elements
// in the same order
func IsEqualSlices(a []int, b []int) bool {
	incompleteA := Any(a, func(i int) bool {
		return i == -1
	})
	incompleteB := Any(b, func(i int) bool {
		return i == -1
	})
	if incompleteA || incompleteB {
		return false
	}
	return slices.Equal(a, b)
}

// TransposeMatrix rotates the matrix 90 degrees clockwise
func TransposeMatrix(matrix [][]int) [][]int {
	transposed := DuplicateBoard(matrix)
	n := len(transposed)
	// Handle the empty matrix case.
	if n == 0 {
		return matrix
	}

	for _, row := range transposed {
		slices.Reverse(row)
	}

	// For every element in the matrix, assign it to the transposed position.
	for i := range n {
		for j := i + 1; j < n; j++ {
			transposed[i][j], transposed[j][i] = transposed[j][i], transposed[i][j]
		}
	}

	return transposed
}

// Map applies a function to every item in the slice
func Map[T any, U any](slice []T, f func(item T) U) []U {
	mapped := []U{}

	for _, item := range slice {
		mapped = append(mapped, f(item))
	}

	return mapped
}
