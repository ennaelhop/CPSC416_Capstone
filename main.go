package main

import (
	"fmt"
)

func main() {

	// Get number of peers from the user
	var numPeers int
	fmt.Println("Enter number of peers: ")
	fmt.Scanf("%d\n", &numPeers)

	// Create a tracker server with 250 peers
	tracker := new(TrackerServer)

	// Create 250 peers and add them to the tracker
	for i := 0; i < numPeers; i++ {
		fmt.Println("Create peer ", i)
		peer := Peer{ID: i, Capacity: 10}
		tracker.Peers = tracker.UpdatePeers(&peer)
	}

	// Create client peer
	client := Peer{Capacity: 10}

	// Get input from the user
	var from, to, targetNumTrailingZeros int

	fmt.Println("Enter from, to, and target number of trailing zeros: ")
	fmt.Scanf("%d %d %d", &from, &to, &targetNumTrailingZeros)

	fmt.Println("Your inputs are: ", from, to, "and", targetNumTrailingZeros)

	fmt.Println("Job inititation request started by client with values: ", from, to, targetNumTrailingZeros)

	// Start a job on the client peer
	client.InitJob(from, to, targetNumTrailingZeros)

}
