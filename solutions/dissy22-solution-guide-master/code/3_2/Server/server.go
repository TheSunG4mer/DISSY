package Server

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	listener       net.Listener
	connections    []net.Conn
	Port           int
	BroadcastCount int
}

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

func (s *Server) broadcast(outbound chan string) {
	for {
		msg_fmt := <-outbound
		for _, conn := range s.connections {
			conn.Write([]byte(msg_fmt))
		}
		s.BroadcastCount++
	}
}

func NewServer() Server {
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

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println("Error is:", err)
	}

	return Server{
		listener: l,
		Port:     l.Addr().(*net.TCPAddr).Port,
	}
}

func (s *Server) Start() {
	defer s.listener.Close()

	outbound := make(chan string)
	go s.broadcast(outbound)

	for {
		fmt.Println("Listening for connection on port ...", s.Port)
		conn, _ := s.listener.Accept()
		s.connections = append(s.connections, conn)
		fmt.Println("Got a connection from ", conn.RemoteAddr().String())

		go handleConnection(conn, outbound)
	}
}
