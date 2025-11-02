package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 1. Get username from the user.
	fmt.Print("Please enter your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	username, err := nameReader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read username:", err)
	}
	username = strings.TrimSpace(username)

	// 2. Send the username to the server as the first message.
	fmt.Fprintln(conn, username)

	// Start a goroutine to continuously read messages from the server
	// and print them to the console.
	go receiveMessages(conn)

	// Use the main goroutine to read input from the user's keyboard
	// and send it to the server.
	fmt.Println("You are now connected. Type a message and press Enter.")
	for {
		msg, _ := nameReader.ReadString('\n')
		text := strings.TrimSpace(msg)
		_, err := fmt.Fprintln(conn, text)
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}

}

// receiveMessages runs in a separate goroutine to handle incoming messages.
func receiveMessages(conn net.Conn) {
	// Copy from the connection's reader to standard output (the console).
	// This will block until the connection is closed or an error occurs.
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Println("Connection lost. ", err)
	}
}
