package main

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
