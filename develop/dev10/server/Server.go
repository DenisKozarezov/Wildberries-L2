package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func HandleConnection(conn net.Conn) {
	log.Printf("Client %s has established a connection to a server.", conn.LocalAddr().String())

	str := "Hello from a server!"

	if _, err := conn.Write([]byte(str)); err != nil {
		log.Printf("Could not send a message to the client: %s", err)
	}

	go func() {
		buffer := make([]byte, 1024)
		for {
			bytesCount, _ := conn.Read(buffer)

			if bytesCount > 0 {
				io.WriteString(os.Stdout, string(buffer)+"\n")
			}
		}
	}()

	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	conn.Close()

	log.Printf("Connection with client %s has been closed due to timeout.", conn.LocalAddr().String())
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	log.Println("Server is listening...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		go HandleConnection(conn)
	}
}
