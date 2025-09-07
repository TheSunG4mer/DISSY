package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	// "time"
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
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ending session with " + port)
			return
		} else {
			fmt.Print("From " + port + ": " + string(msg))
			c <- port + ":" + string(msg)
		}
	}
}

func main() {
	c := make(chan string)
	go broadcast(c)

	// rand.Seed(time.Now().UnixNano())
	port := strconv.Itoa(rand.Intn(9000-1000+1) + 1000)
	// port := "9000" // Use a fixed port
	fmt.Println("Now listening on port", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer ln.Close()
	var new_port string
	for {

		fmt.Println("Listening for connection on port", port, "...")
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		new_port = conn.RemoteAddr().String()
		fmt.Println("Got a connection on port", new_port)
		connections = append(connections, conn)
		go handleConnection(conn, new_port, c)
	}
}
