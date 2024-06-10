package main

import "fmt"

type TrackerServer struct {
	// List of known peers
	Peers []*Peer
}

// Returns a list of peers to the peer requesting it
// TODO: As a nice to have, we can simulate the communication between peers and the tracker
func (ts *TrackerServer) GetPeers() []*Peer {
	fmt.Println("Get peers")
	return ts.Peers
}

// Update the list of peers, used by peers to ping the tracker
func (ts *TrackerServer) UpdatePeers(peer *Peer) []*Peer {
	fmt.Println("Update peers")
	return append(ts.Peers, peer)
}
