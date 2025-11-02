package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	// establishing server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Listening on port 8080 \n")
	// establishing connection to the server

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// ip address fetching
		remoteAddr := conn.RemoteAddr().String()
		host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		log.Printf("Error splitting host and port: %v\n", err)
		return
	}
	fmt.Printf("Client IP address: %s\n", host)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	// Get the remote address (client's IP and port)
	
	for {
		
		n, err := conn.Read(buff)	// returns no. of byte
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Recieved:%s", string(buff[:n])) 
		data := strings.ToUpper(string(buff[:n]))
		_, err = conn.Write([]byte(data))
		if err != nil {
			log.Println(err)
			return
		}
		conn.SetDeadline(time.Now().Add(10 * time.Second))
	}
}
