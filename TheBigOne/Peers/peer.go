package peers

type Peer struct {
	ID              string
	Message_counter int
	Ledger          *Ledger
	IP              string
	Port            int
}

func MakePeer(name string, ip string, port int) *Peer {
	new_peer := new(Peer)
	new_peer.ID = name
	new_peer.Message_counter = 0
	new_peer.Ledger = MakeLedger()
	new_peer.IP = ip
	new_peer.Port = port
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
