package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func sendFile(filename string) {
	conn, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = fmt.Fprintln(conn, filename)
	if err != nil {
		log.Println("Error sending filename:", err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send file contents
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file:", err)
		return
	}

	fmt.Printf("File %s sent successfully.\n", filename)
}
