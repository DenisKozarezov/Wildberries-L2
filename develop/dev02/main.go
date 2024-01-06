package main

import (
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
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
