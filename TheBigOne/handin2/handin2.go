package main

import (
	"Local"
	"Network"
	"fmt"
	"strconv"
	"time"
)

func main() {
	IDs := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	peerList := [10]*Network.Peer{}
	for i, name := range IDs {
		peerList[i] = Network.MakePeer(name, "127.0.0.1", 1001+i)
	}

	clientList := [5]*Local.Client{}
	names := [5]string{}
	for i := 0; i < 5; i++ {
		c, err := Local.MakeClient(200)
		if err != nil {
			panic(err)
		}
		clientList[i] = c
		peerList[0].Ledger.AddParticipant(c.PublicKey)
		fmt.Println(c.PublicKey)
		names[i] = c.PublicKey
	}

	for i, p := range peerList {
		if i == 0 {
			go p.Connect("127.0.0.1", 1000) // should be empty
			time.Sleep(100 * time.Millisecond)
			continue
		}
		port := 1001 // + rand.IntN(i)
		fmt.Printf("Adding peer %v to the network. Connecting on port %v.\n", i+1, port)
		go p.Connect("127.0.0.1", port)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond)

	from := [10]int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3}
	to := [10]int{1, 2, 3, 4, 2, 3, 4, 3, 4, 4}

	fmt.Printf("Peer 1 has %v connections\n", len(peerList[0].OtherPeers))

	for i, peer := range peerList {
		fmt.Printf("Peer %v has %v conns.\n", IDs[i], len(peer.Connections))
	}

	for counter := 0; counter < 10; counter++ {
		for money, p := range peerList {
			tx := Local.MakeTransaction(strconv.Itoa(counter+10*money), names[from[money]], names[to[money]], money+1)
			err := p.FloodSignedTransaction(tx, clientList[from[money]])
			if err != nil {
				fmt.Println(money)
				panic(err)
			}
		}
	}

	time.Sleep(100 * time.Millisecond)

	// Print different

	var index int = 0
	albertBalance, _ := peerList[index].GetLedger().GetBalance(clientList[0].PublicKey)
	bobBalance, _ := peerList[index].GetLedger().GetBalance(clientList[1].PublicKey)
	cathrineBalance, _ := peerList[index].GetLedger().GetBalance(clientList[2].PublicKey)
	danielleBalance, _ := peerList[index].GetLedger().GetBalance(clientList[3].PublicKey)
	eliotBalance, _ := peerList[index].GetLedger().GetBalance(clientList[4].PublicKey)
	fmt.Printf("Balances are:\n\t1: %v\n\t2: %v\n\t3: %v\n\t4: %v\n\t5: %v\n",
		albertBalance, bobBalance, cathrineBalance, danielleBalance, eliotBalance)

	index = 9
	albertBalance, _ = peerList[index].GetLedger().GetBalance(clientList[0].PublicKey)
	bobBalance, _ = peerList[index].GetLedger().GetBalance(clientList[1].PublicKey)
	cathrineBalance, _ = peerList[index].GetLedger().GetBalance(clientList[2].PublicKey)
	danielleBalance, _ = peerList[index].GetLedger().GetBalance(clientList[3].PublicKey)
	eliotBalance, _ = peerList[index].GetLedger().GetBalance(clientList[4].PublicKey)
	fmt.Printf("Balances are:\n\t1: %v\n\t2: %v\n\t3: %v\n\t4: %v\n\t5: %v\n",
		albertBalance, bobBalance, cathrineBalance, danielleBalance, eliotBalance)
	// Should give:
	//  Balances are:
	//		Albert: -100
	//		Bob: -170
	//		Cathrine: -100
	//		Danielle: 70
	//		Eliot: 300
}
