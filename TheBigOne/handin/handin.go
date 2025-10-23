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
	names := [5]string{"Albert", "Bob", "Cathrine", "Danielle", "Eliot"}
	// clientList := [5]*Local.Client{}
	// names := [5]string{}
	// for i := 0; i < 5; i++ {
	// 	c, err := Local.MakeClient(10)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	clientList[i] = c
	// 	fmt.Println(c.PublicKey)
	// 	names[i] = c.PublicKey
	// }
	// for i := 0; i < 5; i++ {
	// 	peerList[0].Ledger.AddParticipant(names[i])
	// 	fmt.Println(names[i])
	// }

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
	time.Sleep(500 * time.Millisecond)

	from := [10]int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3}
	to := [10]int{1, 2, 3, 4, 2, 3, 4, 3, 4, 4}

	fmt.Printf("Peer 1 has %v connections\n", len(peerList[0].OtherPeers))

	for _, peer := range peerList {
		// for _, conn := range peer.Connections {
		// 	fmt.Print("(", conn.LocalAddr(), ",", conn.RemoteAddr(), "), ")
		// }
		fmt.Printf("Number of conns: %v\n", len(peer.Connections))
	}

	for counter := 0; counter < 2; counter++ {
		for i, p := range peerList {
			tx := Local.MakeTransaction(strconv.Itoa(counter+10*i), names[from[i]], names[to[i]], i+1)
			// go p.FloodTransaction(tx)
			p.FloodTransaction(tx)
			// if err != nil {
			// 	fmt.Println(i)
			// 	panic(err)
			// }
			// time.Sleep(1000 * time.Millisecond)
		}
	}

	time.Sleep(10 * time.Millisecond)
	var index int = 0
	albertBalance, _ := peerList[index].GetLedger().GetBalance(names[0])
	bobBalance, _ := peerList[index].GetLedger().GetBalance(names[1])
	cathrineBalance, _ := peerList[index].GetLedger().GetBalance(names[2])
	danielleBalance, _ := peerList[index].GetLedger().GetBalance(names[3])
	eliotBalance, _ := peerList[index].GetLedger().GetBalance(names[4])
	fmt.Printf("Balances are:\n\t1: %v\n\t2: %v\n\t3: %v\n\t4: %v\n\t5: %v\n",
		albertBalance, bobBalance, cathrineBalance, danielleBalance, eliotBalance)

	index = 9
	albertBalance, _ = peerList[index].GetLedger().GetBalance(names[0])
	bobBalance, _ = peerList[index].GetLedger().GetBalance(names[1])
	cathrineBalance, _ = peerList[index].GetLedger().GetBalance(names[2])
	danielleBalance, _ = peerList[index].GetLedger().GetBalance(names[3])
	eliotBalance, _ = peerList[index].GetLedger().GetBalance(names[4])
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
