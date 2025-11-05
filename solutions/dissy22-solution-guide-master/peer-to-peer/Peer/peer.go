package Peer

import (
	"account"
	"encoding/gob"
	"encoding/json"
	"net"
	"strconv"
)

type MessageType string

const (
	GetPeerList         MessageType = "GetPeerList"
	GetPeerListResponse MessageType = "GetPeerListResponse"
	Join                MessageType = "Join"
	Transaction         MessageType = "Transaction"
)

type Message struct {
	Type MessageType
	Data []byte
}

type IP struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type Peer struct {
	IP                IP
	Peers             []IP
	Connections       []*gob.Encoder
	MessageQueue      chan Message
	Ledger            *account.Ledger
	ManualConnections bool //set to true to avoid fully connected network
}

func (p *Peer) Connect(addr string, port int) {
	p.MessageQueue = make(chan Message)
	p.Ledger = account.MakeLedger()

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return
	}

	myAddr := "localhost"

	// Print IP and port of client
	addresses, _ := net.InterfaceAddrs()
	for _, address := range addresses {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.To4()[0] != 169 {
				myAddr = ipnet.IP.String()
				break
			}
		}
	}

	/* TODO: Only works when local */
	p.IP = IP{Addr: myAddr, Port: ln.Addr().(*net.TCPAddr).Port}

	go p.Listen(ln)
	go p.HandleMessages()

	conn, err := net.Dial("tcp", addr+":"+strconv.Itoa(port))

	if err == nil {
		/* Success, setup connection */

		p.HandleConnection(conn)
		data, err := json.Marshal(&p.IP)
		if err != nil {
			return
		}
		p.Peers = append(p.Peers, IP{Addr: addr, Port: port})
		p.FloodMessage(Message{Type: "GetPeerList", Data: data})
	} else {
		/* Failure, create new network */

	}
}

// Make Connection
// Join IP -> ADD IP

func (p *Peer) FloodMessage(msg Message) {
	for _, encoder := range p.Connections {
		encoder.Encode(&msg)
	}
}

func (p *Peer) HandleConnection(conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	p.Connections = append(p.Connections, encoder)
	go p.ReceiveMessages(conn)
}

func (p *Peer) ReceiveMessages(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			/* Error decoding msg */
			continue
		}

		p.MessageQueue <- msg
	}
}

func (p *Peer) HandleMessages() {
	for {
		msg := <-p.MessageQueue
		switch msg.Type {
		case GetPeerList:
			var ip IP
			err := json.Unmarshal(msg.Data, &ip)
			if err != nil {
				return
			}

			data, err := json.Marshal(&p.Peers)
			if err != nil {
				return
			}

			p.Peers = append(p.Peers, ip)

			p.FloodMessage(Message{Type: GetPeerListResponse, Data: data})

		case GetPeerListResponse:
			var peers []IP
			err := json.Unmarshal(msg.Data, &peers)
			if err != nil {
				return
			}

			for _, ip := range peers {
				incorrect := false
				for _, connected_ip := range p.Peers {
					if ip.Port == connected_ip.Port {
						incorrect = true
						break
					}
				}
				if incorrect || ip.Port == p.IP.Port {
					continue
				}

				p.Peers = append(p.Peers, ip)
				if p.ManualConnections {
					continue
				}
				conn, err := net.Dial("tcp", ip.Addr+":"+strconv.Itoa(ip.Port))
				if err != nil {
					return
				}

				p.HandleConnection(conn)
			}

			data, err := json.Marshal(&p.IP)
			if err != nil {
				return
			}

			p.FloodMessage(Message{Type: Join, Data: data})

		case Join:
			var ip IP
			err := json.Unmarshal(msg.Data, &ip)
			if err != nil {
				return
			}
			should_add := true
			for _, connected_ip := range p.Peers {
				if ip.Port == connected_ip.Port {
					should_add = false
					break
				}
			}
			if should_add {
				p.Peers = append(p.Peers, ip)
			}

		case Transaction:
			var tx account.Transaction
			err := json.Unmarshal(msg.Data, &tx)
			if err != nil {
				return
			}
			p.ExecuteTransaction(&tx)

		default:

		}
	}
}

func (p *Peer) Listen(ln net.Listener) {
	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		p.HandleConnection(conn)
	}
}

func (p *Peer) FloodTransaction(tx *account.Transaction) {
	data, err := json.Marshal(tx)
	if err != nil {
		return
	}

	p.FloodMessage(Message{Type: "Transaction", Data: data})
	// p.ExecuteTransaction(tx)
}

func (p *Peer) ExecuteTransaction(tx *account.Transaction) {
	success := p.Ledger.Transaction(tx)
	if success {
		p.FloodTransaction(tx)
	}
}
