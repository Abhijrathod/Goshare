package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	// Start the peer server
	go startPeerServer()

	// Keep the program running
	for {
		var command string
		fmt.Print("Enter command (send <filename> or exit): ")
		fmt.Scan(&command)

		if command == "exit" {
			break
		} else if command == "send" {
			var filename string
			fmt.Scan(&filename)
			sendFile(filename)
		} else {
			fmt.Println("Invalid command. Use 'send <filename>' or 'exit'.")
		}
	}

	fmt.Println("Exiting...")
}

func startPeerServer() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to:", conn.RemoteAddr())

	var filename string
	_, err := fmt.Fscan(conn, &filename)
	if err != nil {
		log.Println("Error reading filename:", err)
		return
	}
	fmt.Printf("Receiving file: %s\n", filename)
	receiveFile(conn, filename) // This should call the receiveFile function defined in file.go
}
