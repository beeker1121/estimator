package utils

// SliceContains checks if a slice contains a given value.
func SliceContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}
