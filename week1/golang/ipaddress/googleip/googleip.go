package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.LookupHost("google.com")
	for indx, addr := range addrs {
		fmt.Println("Address number ", indx, ": ", addr)
	}
}
