package peers

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (p *Peer) Connect(addr string, port int) error {

	go p.AcceptNewConnection()
	go p.AcceptTransactions()
	go p.AcceptNewConnections()

	var dial_up_addr string = addr + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", dial_up_addr)

	if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
		fmt.Printf("Starting new network on %v:%v\n", p.IP, p.Port)
		return p.ListenForNewPeers()
	} else if err != nil {
		return err
	}

	// Initialize starting procedure

	SendString(conn, "Please Gib Contacts!")

	// Recieve everything:
	recievedLedger := false
	numberOfContacts := 0
	targetNumberOfContacts := -1

	for { // get all information one by one
		// and break once they have everything
		m, err := RecieveMessage(conn)
		if err != nil {
			return err
		}
		switch m.Type {
		case MessageString:
			targetNumberOfContacts, err = strconv.Atoi(DemarshalToString(m.Content))
			if err != nil {
				return err
			}
		case MessagePeerInfo:
			p.AddPeer(DemarshalToPeerInfo(m.Content))
			numberOfContacts += 1
		case MessageLedger:
			p.SetLedger(DemarshalToLedger(m.Content))
			recievedLedger = true
		}

		if recievedLedger && numberOfContacts == targetNumberOfContacts {
			break
		}
	}
	conn.Close()

	// Connect to every peer:
	for _, pi := range p.OtherPeers {
		conn, err = net.Dial("tcp", pi.GenerateAddressString())
		if err != nil {
			return err
		}
		SendPeerInfo(conn, p.GeneratePeerInfo())
		p.AddConnection(conn)
		p.ListenToPort(conn)
	}
	return p.ListenForNewPeers()

}

func (p *Peer) ListenForNewPeers() error {
	// Start listening for new peer:
	ln, err := net.Listen("tcp", p.GeneratePeerInfo().GenerateAddressString())
	if err != nil {
		return err
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		p.AddConnection(conn)
		go p.ListenToPort(conn)
	}
}
