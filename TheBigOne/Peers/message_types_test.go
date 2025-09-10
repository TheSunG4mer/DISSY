package peers_test

import (
	peers "P2P_block_chain/Peers"
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
			rawMessage := peers.MarshalStringToMessage(tt.original_input)
			var got string
			if rawMessage.Type == peers.MessageString {
				got = peers.DemarshalToString(rawMessage.Content)
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
		t    *peers.Transaction
		want *peers.Transaction
	}{
		{name: "Normal Transaction", t: peers.MakeTransaction("123", "Bob", "Mike", 10), want: peers.MakeTransaction("123", "Bob", "Mike", 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawMessage := peers.MarshalTransactionToMessage(tt.t)
			var got *peers.Transaction
			if rawMessage.Type == peers.MessageTransaction {
				got = peers.DemarshalToTransaction(rawMessage.Content)
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
	p := peers.MakePeer("Alice", "127.0.0.1", 1234)
	pi := p.GeneratePeerInfo()
	rawMessage := peers.MarshalPeerInfoToMessage(pi)
	if rawMessage.Type != peers.MessagePeerInfo {
		t.Errorf("Did not detect PeerInfo")
		return
	}
	got := peers.DemarshalToPeerInfo(rawMessage.Content)
	if got.ID != pi.ID || got.IP != pi.IP || got.Port != pi.Port {
		t.Errorf("Did not get the right peers back")
	}
}

func TestDemarshalToLedger(t *testing.T) {
	L := peers.MakeLedger()
	L.AddParticipant("Mike")
	L.AddParticipant("John")
	trans := peers.MakeTransaction("123", "Mike", "John", 30)
	L.TranferMoney(trans)

	rawMessage := peers.MarshalLedgerToMessage(L)

	if rawMessage.Type != peers.MessageLedger {
		t.Errorf("Did not detect Ledger")
		return
	}
	got := peers.DemarshalToLedger(rawMessage.Content)
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
	peers.SendString(conn, "message")
	p := peers.MakePeer("Mads", "127.0.0.1", 1234)
	peers.SendPeerInfo(conn, p.GeneratePeerInfo())
	peers.SendTransaction(conn, peers.MakeTransaction("123", "Mike", "John", 10))

	L := peers.MakeLedger()
	L.AddParticipant("Mike")
	L.AddParticipant("John")
	trans := peers.MakeTransaction("123", "Mike", "John", 30)
	L.TranferMoney(trans)

	peers.SendLedger(conn, L)

}

func TestSendingMessageOverNetwork(t *testing.T) {
	listener, er := net.Listen("tcp", "127.0.0.1:5555")
	if er != nil {
		t.Errorf("Error during testing")
		return
	}
	defer listener.Close()
	go helperToTestSendingMessageOverNetwork(t)
	conn, er := listener.Accept()
	if er != nil {
		t.Errorf("Error during testing")
		return
	}
	for i := 0; i < 4; i++ {
		m, er := peers.RecieveMessage(conn)
		if er != nil {
			t.Errorf("Error during testing")
			return
		}
		switch m.Type {
		case peers.MessageString:
			s := peers.DemarshalToString(m.Content)
			if s != "message" {
				t.Errorf("Recieved wrong string")
			}
		case peers.MessagePeerInfo:
			pi := peers.DemarshalToPeerInfo(m.Content)
			if pi.ID != "Mads" || pi.IP != "127.0.0.1" || pi.Port != 1234 {
				t.Errorf("Recieved wrong PeerInfo")
			}
		case peers.MessageTransaction:
			trans := peers.DemarshalToTransaction(m.Content)
			if trans.ID != "123" || trans.From != "Mike" ||
				trans.To != "John" || trans.Amount != 10 {
				t.Errorf("Recieved wrong Transaction")
			}
		case peers.MessageLedger:
			l := peers.DemarshalToLedger(m.Content)
			mikeBalance, _ := l.GetBalance("Mike")
			johnBalance, _ := l.GetBalance("John")
			if mikeBalance != -30 || johnBalance != 30 {
				t.Errorf("Recieved wrong Ledger")
			}

		}
	}

}
