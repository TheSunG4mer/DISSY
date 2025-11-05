package main

import (
	"Automated_client"
	"Server"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	s := Server.NewServer()
	go s.Start()

	startSignal := make(chan struct{})
	waitForSignal := func() {
		<-startSignal // blocks until the channel is closed
	}

	waitThenStartClient := func(client Automated_client.AutoClient) {
		waitForSignal()
		client.Start()
	}

	for i := 0; i < 10; i++ {
		c := Automated_client.NewClient(":" + strconv.Itoa(s.Port))
		go waitThenStartClient(c)
	}
	// Start the clients by closing the signal channel
	close(startSignal)
	// Wait for 2 minutes
	duration := 2 * time.Minute
	t1 := time.NewTimer(duration)
	<-t1.C
	// Write the result to a file
	broadcastCount := s.BroadcastCount
	os.WriteFile("throughput_result", []byte("Broadcast "+strconv.Itoa(broadcastCount)+" messages in "+duration.String()), 0644)
	fmt.Println("Broadcast", broadcastCount, "messages.")
}
