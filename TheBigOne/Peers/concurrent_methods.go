package peers

import (
	"net"
	"strconv"
)

func (p *Peer) AcceptNewConnection() error {
	var conn net.Conn
	for {
		conn = <-p.NewConnectionChannel

		p.MasterLock.Lock()
		var numberOfConnections int = len(p.Connections) + 1
		var stringNumberOfConnections string = strconv.Itoa(numberOfConnections)
		err := SendString(conn, stringNumberOfConnections)
		if err != nil {
			return err
		}
		for _, pi := range p.OtherPeers {
			err = SendPeerInfo(conn, pi)
			if err != nil {
				return err
			}
		}
		err = SendPeerInfo(conn, p.GeneratePeerInfo())
		if err != nil {
			return err
		}
		err = SendLedger(conn, p.Ledger)
		if err != nil {
			return err
		}

		p.MasterLock.Unlock()
	}
}

func (p *Peer) AcceptTransactions() {
	var t *Transaction
	for {
		t = <-p.TransactionChannel
		p.MasterLock.Lock()
		p.Ledger.TranferMoney(t)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) AcceptNewConnections() {
	var pi *PeerInfo
	for {
		pi = <-p.PeerInfoChannel
		p.MasterLock.Lock()
		p.AddPeer(pi)
		p.MasterLock.Unlock()
	}
}

func (p *Peer) ListenToPort(conn net.Conn) error {
	for {
		m, err := RecieveMessage(conn)
		if err != nil {
			return err
		}
		switch m.Type {
		case MessageString:
			switch DemarshalToString(m.Content) {
			case "Please Gib Contacts!":
				p.NewConnectionChannel <- conn
			}
		case MessagePeerInfo:
			p.PeerInfoChannel <- DemarshalToPeerInfo(m.Content)
		case MessageTransaction:
			p.TransactionChannel <- DemarshalToTransaction(m.Content)
		case MessageLedger:
			p.MasterLock.Lock()
			p.SetLedger(DemarshalToLedger(m.Content))
			p.MasterLock.Unlock()
		}
	}
}
