package Network

import (
	"Local"
	"Marshalling"
	"fmt"
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

func (p *Peer) FloodSignedTransaction(tx *Local.Transaction, c *Local.Client) error {
	sgn_tx, err := c.SignTransaction(tx)
	if err != nil {
		fmt.Println("Could not sign transaction")
		return err
	}
	p.MasterLock.Lock()
	p.ApplySignedTransaction(sgn_tx)
	p.MasterLock.Unlock()
	for _, conn := range p.Connections {
		err := Marshalling.SendSignedTransaction(conn, sgn_tx)
		if err != nil {
			return err
		}
	}
	return nil
}
