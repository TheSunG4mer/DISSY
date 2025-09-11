package Network

import (
	"Local"
	"Marshalling"
)

func (p *Peer) FloodMessage(m Marshalling.Message) error {
	for _, conn := range p.Connections {
		err := Marshalling.SendMessage(conn, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Peer) FloodTransaction(tx *Local.Transaction) error {
	p.MasterLock.Lock()
	p.Ledger.TranferMoney(tx)
	p.MasterLock.Unlock()
	for _, conn := range p.Connections {
		err := Marshalling.SendTransaction(conn, tx)
		if err != nil {
			return err
		}
	}
	return nil
}
