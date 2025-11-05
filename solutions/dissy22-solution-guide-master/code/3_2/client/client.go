package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func printResponse(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		msg_ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ending session")
			return
		} else {
			fmt.Print("\nResponse: ", string(msg_ret), "Message: ")
		}
	}
}

func client() {
	var input string

	fmt.Print("IP address and port: ")

	fmt.Scanln(&input)
	fmt.Println(input)

	var conn net.Conn
	conn, _ = net.Dial("tcp", strings.TrimSpace(input))
	defer conn.Close()

	go printResponse(conn)

	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Message: ")
		msg, _ := scanner.ReadString('\n')
		if strings.TrimSpace(msg) == "quit" {
			return
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	client()
}
