package utils

// GetKeys gets all keys from a map
func GetKeys[T comparable, R any](m map[T]R) []T {
	keys := []T{}
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
