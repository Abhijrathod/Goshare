package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func receiveFile(conn net.Conn, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, conn)
	if err != nil {
		log.Println("Error receiving file:", err)
		return
	}

	fmt.Printf("File %s received successfully.\n", filename)
}
