package iteration

import "strings"

// Repeats a given string the given amount of times
func Repeat(character string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(character)
	}
	return repeated.String()
}
