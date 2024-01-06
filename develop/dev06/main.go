package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func cat(lines *[]string, fields []int, delimeter string, separated bool) ([]string, error) {
	var result []string = make([]string, 0, len(*lines))

	for _, str := range *lines {
		if strings.Contains(str, delimeter) {
			splittedWords := strings.Split(str, delimeter)

			var resultStr string

			for _, field := range fields {
				if field <= 0 {
					return nil, errors.New("Invalid field argument!")
				}

				if field < len(splittedWords) {
					resultStr += splittedWords[field-1] + delimeter + " "
				}
			}

			result = append(result, resultStr)
		} else if !separated {
			result = append(result, str)
		}
	}

	return result, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	split := strings.Split(text, "\n")

	result, err := cat(&split, []int{1, 2, 3}, ":", true)
	if err != nil {
		panic(err)
	}

	for _, str := range result {
		fmt.Println(str)
	}
}
