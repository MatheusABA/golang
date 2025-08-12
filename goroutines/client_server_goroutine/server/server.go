package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// go routine to handle each connection, removing will handle one client at a time
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	for {
		_, err := io.WriteString(connection, "Hello from server!\n")
		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
		time.Sleep(time.Second)
	}
}
