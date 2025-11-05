package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn, outbound chan string) {
	defer conn.Close()
	otherEnd := conn.RemoteAddr().String()
	reader := bufio.NewReader(conn) // bufio.NewScanner(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ending session with " + otherEnd)
			return
		} else {
			fmt.Print("From " + otherEnd + ": " + string(msg))
			outbound <- otherEnd + " " + msg
		}
	}
}

var connections []net.Conn

func broadcast(outbound chan string) {
	for {
		msg_fmt := <-outbound
		for _, conn := range connections {
			conn.Write([]byte(msg_fmt))
		}
	}
}

func server() {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error is:", err)
	}
	for _, address := range addresses {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Non-loopback IP is: ", ipnet.IP.String())
			}
		}
		if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Loopback IP is: ", ipnet.IP.String())
			}
		}
	}

	outbound := make(chan string)
	go broadcast(outbound)

	l, err := net.Listen("tcp", ":0") // 18081
	if err != nil {
		fmt.Println("Error is:", err)
	}
	defer l.Close()
	for {		
		fmt.Println("Listening for connection on port ...", l.Addr().(*net.TCPAddr).Port)
		conn, _ := l.Accept()
		connections = append(connections, conn)
		fmt.Println("Got a connection from ", conn.RemoteAddr().String())

		go handleConnection(conn, outbound)
	}
}

func main() {
	server()
}
