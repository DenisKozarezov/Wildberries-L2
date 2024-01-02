package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Key interface {
	handle(lines []string, arg string) []string
}

type After struct {
	count int
}

func (k *After) handle(lines []string, arg string) []string {
	strLen := len(lines)
	hastable := make(map[string]struct{}, strLen)

	for i, str := range lines {
		if strings.Contains(str, arg) {
			hastable[str] = struct{}{}

			var start int = i
			var end int = i + k.count + 1
			if end > strLen {
				end = strLen
			}
			for j := start; j < end; j++ {
				hastable[lines[j]] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(hastable))
	for key, _ := range hastable {
		result = append(result, key)
	}
	return result
}

type Before struct {
	count int
}

func (k *Before) handle(lines []string, arg string) []string {
	strLen := len(lines)
	hastable := make(map[string]struct{}, strLen)

	for i, str := range lines {
		if strings.Contains(str, arg) {
			hastable[str] = struct{}{}

			var start int = i - k.count
			var end int = i
			if start < 0 {
				start = 0
			}
			for j := start; j < end; j++ {
				hastable[lines[j]] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(hastable))
	for key, _ := range hastable {
		result = append(result, key)
	}
	return result
}

type Context struct {
	count int
}

func (k *Context) handle(lines []string, arg string) []string {
	strLen := len(lines)
	hastable := make(map[string]struct{}, strLen)

	for i, str := range lines {
		if strings.Contains(str, arg) {
			var start int = i - k.count
			var end int = i + k.count + 1
			if start < 0 {
				start = 0
			}
			if end > strLen {
				end = strLen
			}
			for j := start; j < end; j++ {
				hastable[lines[j]] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(hastable))
	for key, _ := range hastable {
		result = append(result, key)
	}
	return result
}

type Count struct{}

func (k *Count) handle(lines []string, arg string) int {
	var count int = 0
	for _, str := range lines {
		if strings.Contains(str, arg) {
			count++
		}
	}

	return count
}

type IgnoreCase struct{}

func (k *IgnoreCase) handle(lines []string, arg string) []string {
	var result []string = make([]string, 0, len(lines))

	for i, str := range lines {
		if strings.ContainsAny(str, arg) {
			result = append(result, lines[i])
		}
	}

	return result
}

type Invert struct{}

func (k *Invert) handle(lines []string, arg string) []string {
	var result []string = make([]string, 0, len(lines))

	for i, str := range lines {
		if !strings.Contains(str, arg) {
			result = append(result, lines[i])
		}
	}

	return result
}

type Fixed struct{}

func (k *Fixed) handle(lines []string, arg string) []string {
	var result []string = make([]string, 0, len(lines))

	for i, str := range lines {
		if str == arg {
			result = append(result, lines[i])
		}
	}

	return result
}

type LinedNum struct{}

func (k *LinedNum) handle(lines []string, arg string) []string {
	var result []string = make([]string, 0, len(lines))

	for i, str := range lines {
		if strings.Contains(str, arg) {
			result = append(result, strconv.Itoa(i)+":"+lines[i])
		}
	}

	return result
}

func Grep(str string, arg string, keys ...Key) []string {
	result := strings.Split(str, "\n")
	for _, key := range keys {
		result = key.handle(result, arg)
	}
	return result
}

func main() {
	result := Grep("Welcome to Linux !\n"+
		"Linux is a free and opensource Operating system that is mostly used by\n"+
		"developers and in production servers for hosting crucial components such as web\n"+
		"and database servers. Linux has also made a name for itself in PCs.\n"+
		"Beginners looking to experiment with Linux can get started with friendlier linux\n"+
		"distributions such as Ubuntu, Mint, Fedora and Elementary OS.", "started", &After{count: 1})

	for _, line := range result {
		fmt.Println(line)
	}
}
