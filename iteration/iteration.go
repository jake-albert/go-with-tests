package iteration

import "strings"

func Repeat(s string, l int) string {
	var builder []string
	for i := 0; i < l; i++ {
		builder = append(builder, s)
	}

	return strings.Join(builder, "")
}
