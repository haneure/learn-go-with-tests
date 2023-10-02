package iteration

import "strings"

func Repeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}

func Contains(haysack, needle string) bool {
	return strings.Contains(haysack, needle)
}

func Split(s, separator string) []string {
	return strings.Split(s, separator)
}

func isSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("%s is same with %s\n", v, b[i])
		if v != b[i] {
			// fmt.Printf("%s not same %s\n", v, b[i])
			return false
		}
	}

	return true
}
