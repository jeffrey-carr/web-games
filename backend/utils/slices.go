package utils

import (
	"math/rand"
	"slices"
)

// Any returns true if any elements in slice return true in the
// provided function
func Any[T any](slice []T, f func(x T) bool) bool {
	return slices.ContainsFunc(slice, f)
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

// DuplicateMatrix creates a deep copy of the provided game board
func DuplicateMatrix[T any](matrix [][]T) [][]T {
	var dupe [][]T
	for _, row := range matrix {
		dupeRow := make([]T, len(row))
		copy(dupeRow, row)
		dupe = append(dupe, dupeRow)
	}

	return dupe
}

// TransposeMatrix rotates the matrix 90 degrees clockwise
func TransposeMatrix[T any](matrix [][]T) [][]T {
	transposed := DuplicateMatrix(matrix)
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

// Filter filters out items that do not pass f
func Filter[T comparable](slice []T, f func(item T) bool) []T {
	validItems := []T{}
	for _, item := range slice {
		if f(item) {
			validItems = append(validItems, item)
		}
	}

	return validItems
}

// GetRandomItem gets a random item from a slice
func GetRandomItem[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}
