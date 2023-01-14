package utils

func StringOr(a string, b string) string {
	if a == "" {
		return b
	} else {
		return a
	}
}

func RemoveEmptyStringsFromSlice(strings []string) []string {
	res := make([]string, len(strings))
	copy(res, strings)

	for i, elem := range res {
		if elem == "" || elem == " " {
			res = Remove(res, i)
		}
	}

	return res
}
