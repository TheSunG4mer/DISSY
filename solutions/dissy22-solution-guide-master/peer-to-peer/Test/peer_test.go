package main

import (
	"Peer"
	"account"
	"strconv"
	"testing"
	"time"
)

func TestConnections(t *testing.T) {
	peers := make([]Peer.Peer, 10)

	// Make network
	peers[0] = Peer.Peer{}
	peers[0].Connect("localhost", 18081) // FAILS

	for i := 1; i < 10; i++ {
		peers[i] = Peer.Peer{}
		peers[i].Connect(peers[i-1].IP.Addr, peers[i-1].IP.Port) // Or connect eveything to first peer
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		numPeers := len(peers[i].Peers)
		if numPeers != 9 {
			t.Errorf("Only %d known peers at peer %d; expected 9", numPeers, peers[i].IP.Port)
		}
	}
}

func TestConsistency(t *testing.T) {
	peers := make([]Peer.Peer, 10)

	// Make network
	peers[0] = Peer.Peer{}
	peers[0].Connect("localhost", 18081) // FAILS

	for i := 1; i < 10; i++ {
		peers[i] = Peer.Peer{}
		peers[i].Connect(peers[i-1].IP.Addr, peers[i-1].IP.Port) // Or connect eveything to first peer
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	peers[0].ExecuteTransaction(&account.Transaction{ID: "0", From: "Alice", To: "Bob", Amount: 100})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "1", From: "Carl", To: "Denis", Amount: 108})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "2", From: "Bob", To: "Denis", Amount: 15})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "3", From: "Alice", To: "Carl", Amount: 420})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "4", From: "Carl", To: "Alice", Amount: 69})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "5", From: "Bob", To: "Carl", Amount: 123})

	time.Sleep(1 * time.Second)

	for _, peer := range peers {
		check := peer.Ledger.Accounts["Denis"] == 123 && peer.Ledger.Accounts["Carl"] == 366 && peer.Ledger.Accounts["Bob"] == -38 && peer.Ledger.Accounts["Alice"] == -451
		if !check {
			t.Errorf("Inconsistent ledger at peer %d", peer.IP.Port)
		}
	}
}

func TestConsistencySparseNetwork(t *testing.T) {
	peers := make([]Peer.Peer, 10)

	// Make network
	peers[0] = Peer.Peer{ManualConnections: true}
	peers[0].Connect("localhost", 18081) // FAILS

	for i := 1; i < 10; i++ {
		peers[i] = Peer.Peer{ManualConnections: true}
		peers[i].Connect(peers[i-1].IP.Addr, peers[i-1].IP.Port) // Or connect eveything to first peer
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		numPeers := len(peers[i].Connections)
		if numPeers > 2 {
			t.Errorf("Too many connections (%d) at peer %d", numPeers, peers[i].IP.Port)
		}
	}

	peers[0].ExecuteTransaction(&account.Transaction{ID: "0", From: "Alice", To: "Bob", Amount: 100})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "1", From: "Carl", To: "Denis", Amount: 108})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "2", From: "Bob", To: "Denis", Amount: 15})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "3", From: "Alice", To: "Carl", Amount: 420})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "4", From: "Carl", To: "Alice", Amount: 69})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "5", From: "Bob", To: "Carl", Amount: 123})

	time.Sleep(1 * time.Second)

	for _, peer := range peers {
		check := peer.Ledger.Accounts["Denis"] == 123 && peer.Ledger.Accounts["Carl"] == 366 && peer.Ledger.Accounts["Bob"] == -38 && peer.Ledger.Accounts["Alice"] == -451
		if !check {
			t.Errorf("Inconsistent ledger at peer %d", peer.IP.Port)
		}
	}
}

func main() {
	peers := make([]Peer.Peer, 10)

	// Make network
	peers[0] = Peer.Peer{}
	peers[0].Connect("localhost", 18081) // FAILS

	for i := 1; i < 10; i++ {
		peers[i] = Peer.Peer{}
		peers[i].Connect(peers[i-1].IP.Addr, peers[i-1].IP.Port) // Or connect eveything to first peer
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		print("i" + strconv.Itoa(i) + " " + strconv.Itoa(peers[i].IP.Port) + ":")
		for _, ip := range peers[i].Peers {
			print(strconv.Itoa(ip.Port) + " ")
		}
		println()
	}

	peers[0].ExecuteTransaction(&account.Transaction{ID: "0", From: "Alice", To: "Bob", Amount: 100})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "1", From: "Carl", To: "Denis", Amount: 108})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "2", From: "Bob", To: "Denis", Amount: 15})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "3", From: "Alice", To: "Carl", Amount: 420})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "4", From: "Carl", To: "Alice", Amount: 69})

	peers[0].ExecuteTransaction(&account.Transaction{ID: "5", From: "Bob", To: "Carl", Amount: 123})

	time.Sleep(1 * time.Second)

	for _, peer := range peers {
		println("Ledger status at " + strconv.Itoa(peer.IP.Port))
		for x := range peer.Ledger.Accounts {
			println(x + " " + strconv.Itoa(peer.Ledger.Accounts[x]))
		}
	}
}
