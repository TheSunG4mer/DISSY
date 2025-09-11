package main

import (
	"Local"
	"Network"
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func main() {
	IDs := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	peerList := [10]*Network.Peer{}
	for i, name := range IDs {
		peerList[i] = Network.MakePeer(name, "127.0.0.1", 1001+i)
	}

	peerList[0].Ledger.AddParticipant("Albert")
	peerList[0].Ledger.AddParticipant("Bob")
	peerList[0].Ledger.AddParticipant("Cathrine")
	peerList[0].Ledger.AddParticipant("Danielle")
	peerList[0].Ledger.AddParticipant("Eliot")

	for i, p := range peerList {
		if i == 0 {
			go p.Connect("127.0.0.1", 1000) // should be empty
			time.Sleep(10 * time.Millisecond)
			continue
		}
		port := 1001 + rand.IntN(i)
		fmt.Printf("Adding peer %v to the network. Connecting on port %v.\n", i+1, port)
		go p.Connect("127.0.0.1", port)
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond)

	from := [10]string{"Albert", "Albert", "Albert", "Albert", "Bob", "Bob", "Bob", "Cathrine", "Cathrine", "Danielle"}
	to := [10]string{"Bob", "Cathrine", "Danielle", "Eliot", "Cathrine", "Danielle", "Eliot", "Danielle", "Eliot", "Eliot"}

	fmt.Printf("Peer 1 has %v connections\n", len(peerList[0].OtherPeers))

	for counter := 0; counter < 10; counter++ {
		for i, p := range peerList {
			tx := Local.MakeTransaction(strconv.Itoa(counter+10*i), from[i], to[i], i+1)
			go p.FloodTransaction(tx)
			// time.Sleep(100 * time.Millisecond)
		}
	}

	time.Sleep(10 * time.Millisecond)
	var index int = 0
	albertBalance, _ := peerList[index].GetLedger().GetBalance("Albert")
	bobBalance, _ := peerList[index].GetLedger().GetBalance("Bob")
	cathrineBalance, _ := peerList[index].GetLedger().GetBalance("Cathrine")
	danielleBalance, _ := peerList[index].GetLedger().GetBalance("Danielle")
	eliotBalance, _ := peerList[index].GetLedger().GetBalance("Eliot")
	fmt.Printf("Balances are:\n\tAlbert: %v\n\tBob: %v\n\tCathrine: %v\n\tDanielle: %v\n\tEliot: %v\n",
		albertBalance, bobBalance, cathrineBalance, danielleBalance, eliotBalance)
	// Should give:
	//  Balances are:
	//		Albert: -100
	//		Bob: -170
	//		Cathrine: -100
	//		Danielle: 70
	//		Eliot: 300
}
