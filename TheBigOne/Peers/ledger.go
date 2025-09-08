package peers

import (
	"fmt"
	"sync"
)

type Ledger struct {
	Accounts map[string]int
	lock     sync.Mutex
}

type UserNotFoundError struct {
	name string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("%v", e.name)
}

func MakeLedger() *Ledger {
	ledger := new(Ledger)
	ledger.Accounts = make(map[string]int)
	return ledger
}

func (l *Ledger) AddParticipant(Name string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.Accounts[Name] = 0
}

func (l *Ledger) TranferMoney(T *Transaction) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	_, from_exists := l.Accounts[T.From]
	if !from_exists {
		return &UserNotFoundError{T.From}
	}
	_, to_exists := l.Accounts[T.To]
	if !to_exists {
		return &UserNotFoundError{T.To}
	}

	l.Accounts[T.From] -= T.Amount
	l.Accounts[T.To] += T.Amount
	return nil
}

func (l *Ledger) GetBalance(name string) (int, error) {
	balance, name_exists := l.Accounts[name]
	if !name_exists {
		return 0, &UserNotFoundError{name}
	}
	return balance, nil
}
