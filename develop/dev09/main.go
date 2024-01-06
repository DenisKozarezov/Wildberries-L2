package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func wget(url string) error {
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		return fmt.Errorf("Could not make a GET-query to %s. Reason: %w", url, err)
	}

	reader := bufio.NewReader(response.Body)
	for {
		string, err := reader.ReadString('\n')
		fmt.Println(string)

		if err == io.EOF {
			break
		}
	}

	return nil
}

func main() {
	err := wget("https://www.wildberries.ru/")
	if err != nil {
		log.Fatalln(err)
	}
}
