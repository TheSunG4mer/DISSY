package main

import (
	"bufio"
	"fmt"
	"net"
)

var connections []net.Conn

// holds active connections
func broadcast(c chan string) {
	for {
		msg := <-c
		for _, conn := range connections {
			conn.Write([]byte(msg))
		}
	}
}

func handleConnection(conn net.Conn, port string, c chan string) {
	defer conn.Close()
	otherEnd := conn.RemoteAddr().String()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ending session with " + otherEnd)
			return
		} else {
			fmt.Print("From " + otherEnd + " to " + port + ": " + string(msg))
			c <- port + ":" + string(msg)
		}
	}
}

func main() {
	c := make(chan string)
	go broadcast(c)

	port := "9000" // Use a fixed port
	fmt.Println("Now listening on port", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer ln.Close()
	for {

		fmt.Println("Listening for connection on port", port, "...")
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		fmt.Println("Got a connection on port", port)
		connections = append(connections, conn)
		go handleConnection(conn, port, c)
	}
}
