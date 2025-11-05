package Automated_client

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type AutoClient struct {
	conn net.Conn
}

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

func NewClient(address string) AutoClient {
	conn, _ := net.Dial("tcp", address)

	go printResponse(conn)

	return AutoClient{conn}
}

func (c *AutoClient) Start() {
	myname := c.conn.LocalAddr().String()
	defer c.conn.Close()

	for i := 0; ; i++ {
		msg := myname + "#" + strconv.Itoa(i) + "\n"
		fmt.Println("Sending msg: ", msg)
		c.conn.Write([]byte(msg))
	}
}
