package main

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
