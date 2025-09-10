package peers

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (p *Peer) Connect(addr string, port int) error {

	var dial_up_addr string = addr + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", dial_up_addr)

	if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
		fmt.Printf("Starting new network on %v:%v\n", p.IP, p.Port)
		return p.initialize_network()
	} else if err != nil {
		return err
	}
	SendMessage(conn, Message{}) // placeholder to get rid of errors.

	return nil
}

func (p *Peer) initialize_network() error {
	panic("unimplemented")
}
