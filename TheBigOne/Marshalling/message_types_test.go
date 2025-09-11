package Marshalling

import (
	"Local"
	"encoding/json"

	// "Network"
	"net"
	"testing"
)

func TestDemarshalToString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		original_input string
		want           string
	}{
		{name: "Number String", original_input: "1234", want: "1234"},
		{name: "Letter String", original_input: "abcd", want: "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawMessage := MarshalStringToMessage(tt.original_input)
			var got string
			if rawMessage.Type == MessageString {
				got = DemarshalToString(rawMessage.Content)
			} else {
				t.Errorf("Did not detect string")
				got = tt.want
			}
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("DemarshalToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemarshalToTransaction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		t    *Local.Transaction
		want *Local.Transaction
	}{
		{name: "Normal Transaction", t: Local.MakeTransaction("123", "Bob", "Mike", 10), want: Local.MakeTransaction("123", "Bob", "Mike", 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawMessage := MarshalTransactionToMessage(tt.t)
			var got *Local.Transaction
			if rawMessage.Type == MessageTransaction {
				got = DemarshalToTransaction(rawMessage.Content)
			} else {
				t.Errorf("Did not detect transaction")
				got = tt.want
			}
			// TODO: update the condition below to compare got with tt.want.
			if tt.want.ID != got.ID || tt.want.From != got.From || tt.want.To != got.To || tt.want.Amount != got.Amount {
				t.Errorf("DemarshalToTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemarshalToPeerInfo(t *testing.T) {
	pi := Local.PeerInfo{ID: "Alice", IP: "127.0.0.1", Port: 1234}
	rawMessage := MarshalPeerInfoToMessage(&pi)
	if rawMessage.Type != MessagePeerInfo {
		t.Errorf("Did not detect PeerInfo")
		return
	}
	got := DemarshalToPeerInfo(rawMessage.Content)
	if got.ID != pi.ID || got.IP != pi.IP || got.Port != pi.Port {
		t.Errorf("Did not get the right peers back")
	}
}

func TestDemarshalToLedger(t *testing.T) {
	L := Local.MakeLedger()
	L.AddParticipant("Mike")
	L.AddParticipant("John")
	trans := Local.MakeTransaction("123", "Mike", "John", 30)
	L.TranferMoney(trans)

	rawMessage := MarshalLedgerToMessage(L)

	if rawMessage.Type != MessageLedger {
		t.Errorf("Did not detect Ledger")
		return
	}
	got := DemarshalToLedger(rawMessage.Content)
	johnNewBalance, _ := got.GetBalance("John")
	mikeNewBalance, _ := got.GetBalance("Mike")
	if johnNewBalance != 30 || mikeNewBalance != -30 {
		t.Errorf("Did not get the right ledger back")
	}
}

func helperToTestSendingMessageOverNetwork(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		t.Errorf("Error during testing")
		return
	}
	SendString(conn, "message")
	t.Log("String sent")
	p := Local.PeerInfo{ID: "Mads", IP: "127.0.0.1", Port: 1234}
	SendPeerInfo(conn, &p)
	t.Log("PeerInfo sent")
	SendTransaction(conn, Local.MakeTransaction("123", "Mike", "John", 10))
	t.Log("Transaction sent")

	L := Local.MakeLedger()
	L.AddParticipant("Mike")
	L.AddParticipant("John")
	trans := Local.MakeTransaction("123", "Mike", "John", 30)
	L.TranferMoney(trans)

	SendLedger(conn, L)
	t.Log("Ledger sent")

}

func TestSendingMessageOverNetwork(t *testing.T) {
	listener, er := net.Listen("tcp", "127.0.0.1:5555")
	if er != nil {
		t.Errorf("Error during testing")
		return
	}
	defer listener.Close()
	go helperToTestSendingMessageOverNetwork(t)
	t.Log("Started client")
	conn, er := listener.Accept()
	if er != nil {
		t.Errorf("Error during testing")
		return
	}
	t.Log("Connected to client")
	dec := json.NewDecoder(conn)
	for i := 0; i < 4; i++ {
		m, er := RecieveMessage(dec)
		if er != nil {
			t.Errorf("Error during testing")
			return
		}
		switch m.Type {
		case MessageString:
			s := DemarshalToString(m.Content)
			if s != "message" {
				t.Errorf("Recieved wrong string")
			}
			t.Log("Recieved String")
		case MessagePeerInfo:
			pi := DemarshalToPeerInfo(m.Content)
			if pi.ID != "Mads" || pi.IP != "127.0.0.1" || pi.Port != 1234 {
				t.Errorf("Recieved wrong PeerInfo")
			}
			t.Log("Recieved PeerInfo")
		case MessageTransaction:
			trans := DemarshalToTransaction(m.Content)
			if trans.ID != "123" || trans.From != "Mike" ||
				trans.To != "John" || trans.Amount != 10 {
				t.Errorf("Recieved wrong Transaction")
			}
			t.Log("Recieved Transaction")
		case MessageLedger:
			l := DemarshalToLedger(m.Content)
			mikeBalance, _ := l.GetBalance("Mike")
			johnBalance, _ := l.GetBalance("John")
			if mikeBalance != -30 || johnBalance != 30 {
				t.Errorf("Recieved wrong Ledger")
			}
			t.Log("Recieved Ledger")
		}
	}
}
