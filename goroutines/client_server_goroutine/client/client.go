package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	defer connection.Close()

	mustCopy(os.Stdout, connection)

}

func mustCopy(ds io.Writer, src io.Reader) {
	if _, err := io.Copy(ds, src); err != nil {
		log.Fatal("Error copying data:", err)
	}
	log.Println("Data copied successfully")
}
