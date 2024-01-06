package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
}

func (c *Client) Telnet(ip string, port int, timeoutInSec int) error {
	log.Printf("Attempting to telnet IP = %s Port = %d\n", ip, port)

	address := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	socket, err := net.DialTimeout("tcp", address, time.Second*time.Duration(timeoutInSec))

	if err != nil {
		return fmt.Errorf("Unable to telnet: %w", err)
	}

	log.Println("Connection is established.")

	var connectionClosed bool = false
	go func() {
		buffer := make([]byte, 1024)
		for {
			bytesCount, err := socket.Read(buffer)
			if e, ok := err.(net.Error); ok && e.Timeout() {
				log.Printf("Connection with client %s has been closed due to timeout.", socket.LocalAddr().String())
				connectionClosed = true
				return
			}

			if bytesCount > 0 {
				io.WriteString(os.Stdout, string(buffer)+"\n")
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if connectionClosed {
			return nil
		}

		if scanner.Scan() {
			send := scanner.Bytes()

			_, err := socket.Write(send)
			if err != nil {
				log.Println("Could not send a message to the server:", err)
			}
		}
	}
}

func main() {
	client := &Client{}

	err := client.Telnet("127.0.0.1", 8080, 10)

	if err != nil {
		panic(err)
	}
}
