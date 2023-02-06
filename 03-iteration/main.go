package iteration

import "strings"

func Repeat(c rune, times int) string {
	sb := new(strings.Builder)
	for i := 0; i < times; i++ {
		sb.WriteRune(c)
	}
	return sb.String()
}

func RepeatByConcatenation(c rune, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += string(c)
	}
	return result
}
