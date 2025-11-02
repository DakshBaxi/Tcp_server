package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func main() {
	// establishing server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("Listening on port 8080 \n")
	// establishing connection to the server
	// Using sync.Map to not deal with concurrency slice/map issues
	var connMap = &sync.Map{}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		// storing each conn in map
		id := uuid.New().String()
		connMap.Store(id, conn)

		// ip address fetching
		remoteAddr := conn.RemoteAddr().String()
		host, _, err := net.SplitHostPort(remoteAddr)
		if err != nil {
			log.Printf("Error splitting host and port: %v\n", err)
			return
		}
		fmt.Printf("Client IP address: %s\n", host)
		go handleConn(conn, id, connMap)
	}
}

func handleConn(conn net.Conn, id string, connMap *sync.Map) {
	// First, read the username from the connection
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Failed to read username:", err)
		conn.Close()
		connMap.Delete(id)
		return
	}
	username = strings.TrimSpace(username)

	// Announce that a new user has joined
	broadcastMessage(fmt.Sprintf("--- %s has joined the chat ---\n", username), connMap, nil)

	defer func() {
		conn.Close()
		connMap.Delete(id)
		// Announce that the user has left
		broadcastMessage(fmt.Sprintf("--- %s has left the chat ---\n", username), connMap, nil)
	}()

	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// This error usually means the client has disconnected.
			// The defer function will handle cleanup and announcement.
			return
		}

		// Prepend username to the message and broadcast it
		message := fmt.Sprintf("%s: %s", username, userInput)
		broadcastMessage(message, connMap, conn)
	}
}

// broadcastMessage sends a message to all clients in the connMap.
// If sender is not nil, it skips sending the message back to the original sender.
func broadcastMessage(message string, connMap *sync.Map, sender net.Conn) {
	connMap.Range(func(key, value interface{}) bool {
		if targetConn, ok := value.(net.Conn); ok {
			// Don't send the message back to the sender
			if targetConn != sender {
				_, err := targetConn.Write([]byte(message))
				if err != nil {
					log.Println("Removing disconnected client:", err)
					targetConn.Close()
					connMap.Delete(key)
				}
			}
		}
		return true
	})
}
