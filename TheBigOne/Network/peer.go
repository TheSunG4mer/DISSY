package Network

import (
	"Local"
	"net"
	"sync"
)

type Peer struct {
	ID                   string
	Message_counter      int
	Ledger               *Local.Ledger
	IP                   string
	Port                 int
	NewConnectionChannel chan net.Conn
	TransactionChannel   chan *Local.Transaction
	PeerInfoChannel      chan *Local.PeerInfo
	OtherPeers           []*Local.PeerInfo
	Connections          []net.Conn
	MasterLock           sync.Mutex
}

func MakePeer(name string, ip string, port int) *Peer {
	new_peer := new(Peer)
	new_peer.ID = name
	new_peer.Message_counter = 0
	new_peer.Ledger = Local.MakeLedger()
	new_peer.IP = ip
	new_peer.Port = port

	new_peer.NewConnectionChannel = make(chan net.Conn)
	new_peer.TransactionChannel = make(chan *Local.Transaction)
	new_peer.PeerInfoChannel = make(chan *Local.PeerInfo)

	new_peer.OtherPeers = []*Local.PeerInfo{}
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

func (p *Peer) GetLedger() *Local.Ledger {
	return p.Ledger
}

func (p *Peer) ApplyTransaction(t *Local.Transaction) {
	p.Ledger.TranferMoney(t)
}

func (p *Peer) GeneratePeerInfo() *Local.PeerInfo {
	pi := new(Local.PeerInfo)
	pi.ID = p.ID
	pi.IP = p.IP
	pi.Port = p.Port
	return pi
}

func (p *Peer) SetLedger(l *Local.Ledger) {
	p.Ledger = l
}

func (p *Peer) AddPeer(pi *Local.PeerInfo) {
	p.OtherPeers = append(p.OtherPeers, pi)
}

func (p *Peer) AddConnection(conn net.Conn) {
	p.Connections = append(p.Connections, conn)
}
