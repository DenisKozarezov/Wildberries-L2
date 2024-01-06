package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Key interface {
	sortByKey(lines *[]string)
}

type K struct {
	column int
}

func (k *K) sortByKey(lines *[]string) {
	sort.Slice(*lines, func(i, j int) bool { return (*lines)[i] < (*lines)[k.column] })
}

type N struct{}

func (k *N) sortByKey(lines *[]string) {
	var numbers []int
	for _, str := range *lines {
		number, err := strconv.Atoi(str)
		if err != nil {
			continue
		} else {
			numbers = append(numbers, number)
		}
	}

	sort.Ints(numbers)

	var result []string = make([]string, 0, len(numbers))
	for _, number := range numbers {
		result = append(result, strconv.Itoa(number))
	}
	*lines = result
}

type R struct{}

func (k *R) sortByKey(lines *[]string) {
	sort.Slice(*lines, func(i, j int) bool { return (*lines)[i] > (*lines)[j] })
}

type U struct{}

func (k *U) sortByKey(lines *[]string) {
	hashtable := make(map[string]struct{}, len(*lines))

	for _, str := range *lines {
		_, isFound := hashtable[str]
		if !isFound {
			hashtable[str] = struct{}{}
		}
	}

	result := make([]string, 0, len(hashtable))
	for key, _ := range hashtable {
		result = append(result, key)
	}
	*lines = result
}

func SortFileString(filepath string, keys ...Key) {
	bytes, _ := os.ReadFile(filepath)
	lines := strings.Split(string(bytes), " ")

	for _, key := range keys {
		key.sortByKey(&lines)
	}

	os.WriteFile(filepath, []byte(strings.Join(lines, " ")), 0644)
}

func main() {
	SortFileString("myfile.txt", &U{}, &N{}, &R{}, &K{column: 3})
}
