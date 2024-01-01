package main

import (
	"strings"
	"unicode"
)

/*
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd
"45" => "" (некорректная строка)
"" => ""
"a45" => "" (некорректная строка)
*/

func UnpackString(zippedStr string) string {
	var result string = ""
	var lastChar rune
	for _, r := range zippedStr {
		isDigit := unicode.IsDigit(r)

		if !isDigit {
			result += string(r)
		} else {
			if lastChar == 0 || unicode.IsDigit(lastChar) {
				return ""
			}

			code := (int)(r) - (int)('0')
			result += strings.Repeat(string(lastChar), code-1)
		}
		lastChar = r
	}
	return result
}
