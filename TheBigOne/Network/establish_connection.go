package Network

import (
	"Marshalling"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (p *Peer) Connect(addr string, port int) error {

	go p.AcceptNewConnection()
	go p.AcceptTransactions()
	go p.AcceptNewConnections()
	go p.AcceptSignedTransactions()

	var dial_up_addr string = addr + ":" + strconv.Itoa(port)
	// fmt.Printf("%v is ready to dial\n", p.ID)
	conn, err := net.Dial("tcp", dial_up_addr)

	if err != nil {
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			fmt.Printf("Starting new network on %v:%v\n", p.IP, p.Port)
			return p.ListenForNewPeers()
		}
		return err
	}

	// fmt.Printf("%v got %v as dial up\n", p.ID, conn.RemoteAddr())
	// Initialize starting procedure

	Marshalling.SendString(conn, "Please Gib Contacts!")
	dec := json.NewDecoder(conn)
	// Recieve everything:
	recievedLedger := false
	numberOfContacts := 0
	targetNumberOfContacts := -1

	// fmt.Printf("%v is ready to receive info\n", p.ID)
	for { // get all information one by one
		// and break once they have everything
		m, err := Marshalling.RecieveMessage(dec)
		// fmt.Printf("%v received message of type %v\n", p.ID, m.Type)
		if err != nil {
			return err
		}
		switch m.Type {
		case Marshalling.MessageString:
			targetNumberOfContacts, err = strconv.Atoi(Marshalling.DemarshalToString(m.Content))
			if err != nil {
				return err
			}
		case Marshalling.MessagePeerInfo:
			p.AddPeer(Marshalling.DemarshalToPeerInfo(m.Content))
			numberOfContacts += 1
		case Marshalling.MessageLedger:
			p.SetLedger(Marshalling.DemarshalToLedger(m.Content))
			recievedLedger = true
		}
		// fmt.Printf("%v is checking: Has received ledger: %v. Contacts: %v/%v\n", p.ID, recievedLedger, numberOfContacts, targetNumberOfContacts)
		if recievedLedger && numberOfContacts == targetNumberOfContacts {
			break
		}
	}

	conn.Close()
	// fmt.Printf("%v escaped initialization\n", p.ID)
	// Connect to every peer:
	for _, pi := range p.OtherPeers {
		conn, err = net.Dial("tcp", pi.GenerateAddressString())
		if err != nil {
			return err
		}
		Marshalling.SendPeerInfo(conn, p.GeneratePeerInfo())
		p.AddConnection(conn)
		go p.ListenToPort(conn)
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
	p.AddPeer(p.GeneratePeerInfo())
	fmt.Printf("%v has started to listen to their port\n", p.ID)
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		p.AddConnection(conn)
		// fmt.Printf("%v got connection from %v\n", p.ID, conn.RemoteAddr())
		go p.ListenToPort(conn)
	}
}
