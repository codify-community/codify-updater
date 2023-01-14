package utils

func StringOr(a string, b string) string {
	if a == "" {
		return b
	} else {
		return a
	}
}
