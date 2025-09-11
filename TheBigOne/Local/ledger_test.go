package Local

import (
	"testing"
)

func TestLedger_TranferMoney(t *testing.T) {

	l := MakeLedger()
	l.AddParticipant("John")
	l.AddParticipant("Dave")
	l.AddParticipant("Michael")

	john_balance, er := l.GetBalance("John")
	if john_balance != 0 {
		t.Errorf("Initial balance was not 0")
	}
	if er != nil {
		t.Errorf("Encountered error reading initialized balance")
	}

	_, er = l.GetBalance("Sarah")
	if er == nil {
		t.Errorf("No error encounted reading from non-existing account")
	}

	transaction := MakeTransaction("1", "John", "Dave", 50)
	l.TranferMoney(transaction)

	john_balance, er = l.GetBalance("John")
	if john_balance != -50 {
		t.Errorf("Sender not losing correct amount of money")
	}
	if er != nil {
		t.Errorf("Encountering error when reading balance of sender")
	}

	dave_balance, er := l.GetBalance("Dave")
	if dave_balance != 50 {
		t.Errorf("Receiver not gaining correct amount of money")
	}
	if er != nil {
		t.Errorf("Encountering error when reading balance of receiver")
	}

}
