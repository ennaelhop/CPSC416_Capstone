package main

import (
	"fmt"
	"testing"
)

func BenchmarkTrackerServer_Parallel(b *testing.B) {
	// Create a tracker server with 250 peers
	tracker := new(TrackerServer)

	// Create 250 peers and add them to the tracker
	for i := 0; i < 250; i++ {
		fmt.Println("Create peer ", i)
		peer := Peer{ID: i, Capacity: 10, Peers: tracker.Peers}
		tracker.Peers = tracker.UpdatePeers(&peer)
	}

	// Create client peer
	client := Peer{Capacity: 10, Peers: tracker.Peers}

	setupBenchmark()

	client.InitJob(0, 10000, 6)
}

func BenchmarkTrackerServer_GetPeers_Single(b *testing.B) {
	// Create a tracker server with 250 peers
	tracker := new(TrackerServer)

	// Create 250 peers and add them to the tracker
	for i := 0; i < 250; i++ {
		fmt.Println("Create peer ", i)
		peer := Peer{ID: i, Capacity: 10}
		tracker.Peers = tracker.UpdatePeers(&peer)
	}

	// Create client peer
	client := Peer{Capacity: 10}

	setupBenchmark()

	client.StartJob(0, 10000, 6)
}
