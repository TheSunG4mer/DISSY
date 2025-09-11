package peers

func (p *Peer) FloodMessage(m Message) error {
	for _, conn := range p.Connections {
		err := SendMessage(conn, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Peer) FloodTransaction(tx *Transaction) error {
	for _, conn := range p.Connections {
		err := SendTransaction(conn, tx)
		if err != nil {
			return err
		}
	}
	return nil
}
