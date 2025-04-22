package utils

// Zero returns the zero value of a type
func Zero[T any]() T {
	var item T
	return item
}

// IsZero returns true if the item is the zero-value of its type
func IsZero[T comparable](item T) bool {
	var zero T
	return item == zero
}
