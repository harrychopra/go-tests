package iteration

import "strings"

func Repeat(c rune, times int) string {
	sb := new(strings.Builder)
	for i := 0; i < times; i++ {
		sb.WriteRune(c)
	}
	return sb.String()
}
