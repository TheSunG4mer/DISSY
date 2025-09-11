package peers

import (
	"net"
	"sync"
)

type Peer struct {
	ID                   string
	Message_counter      int
	Ledger               *Ledger
	IP                   string
	Port                 int
	NewConnectionChannel chan net.Conn
	TransactionChannel   chan *Transaction
	PeerInfoChannel      chan *PeerInfo
	OtherPeers           []*PeerInfo
	Connections          []net.Conn
	MasterLock           sync.Mutex
}

func MakePeer(name string, ip string, port int) *Peer {
	new_peer := new(Peer)
	new_peer.ID = name
	new_peer.Message_counter = 0
	new_peer.Ledger = MakeLedger()
	new_peer.IP = ip
	new_peer.Port = port

	new_peer.NewConnectionChannel = make(chan net.Conn)
	new_peer.TransactionChannel = make(chan *Transaction)
	new_peer.PeerInfoChannel = make(chan *PeerInfo)

	new_peer.OtherPeers = []*PeerInfo{}
	new_peer.Connections = []net.Conn{}

	new_peer.MasterLock = sync.Mutex{}

	return new_peer
}

func (p *Peer) GetName() string {
	return p.ID
}

func (p *Peer) GetMessageCounter() int {
	return p.Message_counter
}

func (p *Peer) BumpMessageCounter() int {
	toReturn := p.GetMessageCounter()
	p.Message_counter++
	return toReturn
}

func (p *Peer) GetLedger() *Ledger {
	return p.Ledger
}

func (p *Peer) ApplyTransaction(t *Transaction) {
	p.Ledger.TranferMoney(t)
}

func (p *Peer) GeneratePeerInfo() *PeerInfo {
	pi := new(PeerInfo)
	pi.ID = p.ID
	pi.IP = p.IP
	pi.Port = p.Port
	return pi
}

func (p *Peer) SetLedger(l *Ledger) {
	p.Ledger = l
}

func (p *Peer) AddPeer(pi *PeerInfo) {
	p.OtherPeers = append(p.OtherPeers, pi)
}
