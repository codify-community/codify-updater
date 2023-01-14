package utils

func Has[T comparable](arr []T, val T) (T, bool) {
	for _, i := range arr {
		if i == val {
			return i, true
		}
	}

	return val, false
}

func Remove[Type interface{}](s []Type, i int) []Type {
	return append(s[:i], s[i+1:]...)
}
