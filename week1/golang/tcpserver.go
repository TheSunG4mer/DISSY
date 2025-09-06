package main

import (
	"bufio" ; "fmt" ; "net" ; "strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		} else {
			fmt.Println("From Client:", string(msg))
			titlemsg := strings.Title(msg)
			conn.Write([]byte(titlemsg))
		}
	}
}

func main() {
	fmt.Println("Listening for connection...")
	ln, _ := net.Listen("tcp", ":18081")
	defer ln.Close()
	conn, _ := ln.Accept()
	fmt.Println("Got a connection...")
	handleConnection(conn)
}
