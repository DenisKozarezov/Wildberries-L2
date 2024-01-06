package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode/utf8"
)

type AnagramSet []string
type AnagramsContainer map[string]*AnagramSet

func IsAnagramOfWord(target string, source string) bool {
	var count int = 0
	for _, r := range source {
		contains := strings.ContainsRune(target, r)
		if !contains {
			return false
		} else {
			count++
		}
	}
	return count == utf8.RuneCountInString(source)
}

func FindAnagrams(array *[]string) *AnagramsContainer {
	hashtable := make(map[string]*AnagramSet)

	for i, str := range *array {
		str = strings.ToLower(str)
		(*array)[i] = str
		hashtable[str] = &AnagramSet{}
	}

	for key, set := range hashtable {
		for _, str := range *array {
			if !slices.Contains(*set, key) && IsAnagramOfWord(str, key) {
				*set = append(*set, str)
			}
		}
	}

	result := make(AnagramsContainer, len(hashtable))
	for key, set := range hashtable {
		if len(*set) <= 1 {
			continue
		}

		sort.Strings(*set)
		result[key] = set
	}

	for key1, set1 := range result {
		for key2, set2 := range result {
			if IsAnagramOfWord(key1, key2) && len(*set1) < len(*set2) {
				delete(result, key1)
			}
		}
	}

	return &result
}

func main() {
	var strings []string = []string{
		"тЯпка",
		"пятка",
		"Слиток",
		"листок",
		"пятак",
		"столик",
		"Тяпка",
		"слиток",
		"банан",
		"Банан",
		"банан",
	}

	var strings2 []string = []string{"пятак", "пятка", "тяпка", "листок", "слиток"}

	result := FindAnagrams(&strings)
	for key, set := range *result {
		fmt.Printf("Key: %s, Set: %s\n", key, *set)
	}

	fmt.Println()

	result = FindAnagrams(&strings2)
	for key, set := range *result {
		fmt.Printf("Key: %s, Set: %s\n", key, *set)
	}
}
