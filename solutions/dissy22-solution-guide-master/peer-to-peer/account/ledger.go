package account

import (
	"sync"
)

type Ledger struct {
	Accounts             map[string]int
	PreviousTransactions map[string]bool
	lock                 sync.Mutex
}

func MakeLedger() *Ledger {
	ledger := new(Ledger)
	ledger.Accounts = make(map[string]int)
	ledger.PreviousTransactions = make(map[string]bool)
	return ledger
}

type Transaction struct {
	ID     string
	From   string
	To     string
	Amount int
}

func (l *Ledger) Transaction(t *Transaction) (success bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if !l.PreviousTransactions[t.ID] {
		l.PreviousTransactions[t.ID] = true
		l.Accounts[t.From] -= t.Amount
		l.Accounts[t.To] += t.Amount
		return true
	}
	return false
}
