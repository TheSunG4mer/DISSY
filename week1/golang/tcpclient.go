package main

import (
	"bufio" ; "fmt" ; "net" ; "os" ;"strings"
)

var conn net.Conn

func main() {
	conn, _ = net.Dial("tcp", "127.0.0.1:18081")
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if text == "quit\n" {
			return
		}
		fmt.Fprintf(conn, text)
		msg, err := connReader.ReadString('\n')
		if err != nil {
			return
		}
		// Windows uses \r\n as a return character, Trimspace removes the extra '\r' character
		msg = strings.TrimSpace(msg)
		fmt.Println("From server: " + msg)
	}
}
