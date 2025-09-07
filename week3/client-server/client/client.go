package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var conn net.Conn

func print_incomming() {
	connReader := bufio.NewReader(conn)
	for {
		msg, err := connReader.ReadString('\n')
		if err != nil {
			return
		}
		// Windows uses \r\n as a return character, Trimspace removes the extra '\r' character
		msg = strings.TrimSpace(msg)
		fmt.Println("From server: " + msg)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input ip address and port number (addr:port): ")
	ip_port, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	ip_port = strings.TrimSpace(ip_port)
	conn, _ = net.Dial("tcp", ip_port) //"127.0.0.1:9000"
	defer conn.Close()
	go print_incomming()

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if text == "quit\n" {
			return
		}
		fmt.Fprintf(conn, "%s", text)

	}
}
