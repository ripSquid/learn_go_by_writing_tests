package iteration

import "strings"

func Repeat(character string, times_repeated int) string {
	var result strings.Builder
	for i := 0; i < times_repeated; i++ {
		result.WriteString(character)
	}
	return result.String()
}
