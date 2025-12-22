package cli

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepareData(str)

	return ""
}

func prepareData(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
