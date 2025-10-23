package Network

import (
	"Local"
	"Marshalling"
	"encoding/json"

	// "fmt"
	"net"
	"strconv"
)

func (p *Peer) AcceptNewConnection() error {
	var conn net.Conn
	for {
		conn = <-p.NewConnectionChannel

		p.MasterLock.Lock()

		var numberOfConnections int = len(p.OtherPeers)
		var stringNumberOfConnections string = strconv.Itoa(numberOfConnections)
		err := Marshalling.SendString(conn, stringNumberOfConnections)
		if err != nil {
			return err
		}
		// fmt.Printf("	%v sent string %v to %v\n", p.ID, stringNumberOfConnections, conn.RemoteAddr())

		for _, pi := range p.OtherPeers {
			err = Marshalling.SendPeerInfo(conn, pi)
			if err != nil {
				return err
			}
			// fmt.Printf("	%v sent peer to %v\n", p.ID, conn.RemoteAddr())
		}

		err = Marshalling.SendLedger(conn, p.Ledger)
		if err != nil {
			return err
		}
		// fmt.Printf("	%v sent the ledger to %v\n", p.ID, conn.RemoteAddr())
		p.RemoveConnection(conn)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) AcceptTransactions() {
	var t *Local.Transaction
	for {
		t = <-p.TransactionChannel
		p.MasterLock.Lock()
		p.ApplyTransaction(t)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) AcceptSignedTransactions() {
	var sgn_tx *Local.SignedTransaction
	for {
		sgn_tx = <-p.SignedTransactionChannel
		// fmt.Println("Received signed transaction")
		p.MasterLock.Lock()
		p.ApplySignedTransaction(sgn_tx)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) AcceptNewConnections() {
	var pi *Local.PeerInfo
	for {
		pi = <-p.PeerInfoChannel
		p.MasterLock.Lock()
		p.AddPeer(pi)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) ListenToPort(conn net.Conn) error {

	jsonDecoder := json.NewDecoder(conn)
	for {
		m, err := Marshalling.RecieveMessage(jsonDecoder)
		if err != nil {
			return err
		}
		switch m.Type {
		case Marshalling.MessageString:
			switch Marshalling.DemarshalToString(m.Content) {
			case "Please Gib Contacts!":
				p.NewConnectionChannel <- conn
			}
		case Marshalling.MessagePeerInfo:
			p.PeerInfoChannel <- Marshalling.DemarshalToPeerInfo(m.Content)
		case Marshalling.MessageTransaction:
			p.TransactionChannel <- Marshalling.DemarshalToTransaction(m.Content)
		case Marshalling.MessageSignedTransaction:
			// fmt.Println("Received signed Transaction")
			p.SignedTransactionChannel <- Marshalling.DemarshalToSignedTransaction(m.Content)
		case Marshalling.MessageLedger:
			p.MasterLock.Lock()
			p.SetLedger(Marshalling.DemarshalToLedger(m.Content))
			p.MasterLock.Unlock()
		}
	}
}
