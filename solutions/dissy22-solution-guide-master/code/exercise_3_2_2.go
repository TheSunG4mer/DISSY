package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"strconv"
	"os"
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

	if (len(os.Args) == 1) {
		fmt.Print("IP address and port: ")

		fmt.Scanln(&input)
		fmt.Println(input)
	} else {
		 input = os.Args[1];
	}
	fmt.Print("Input: ", input);

	var conn net.Conn
	conn, _ = net.Dial("tcp", strings.TrimSpace(input))
	defer conn.Close()

	go printResponse(conn)

	myname := conn.LocalAddr().String()

	for i := 0; ; i++ {
		msg := myname + "#" + strconv.Itoa(i) + "\n"
		fmt.Println("Sending msg: ", msg)
		conn.Write([]byte(msg))
	}
}

func main() {
	client()
}
