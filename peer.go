package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type Peer struct {
	// ID of the peer
	ID int
	// The capacity of the peer in terms of number of jobs it can handle
	Capacity int
	// List of known peers
	Peers []*Peer
}

type Block struct {
	// Realistically this block will have many other fields
	// but we are abstracting it into a single integer
	Nonce int
	Hash  string
}

const MAX_SIZE = 25

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// InitJob initiates a job with the given parameters
//
// from: the starting value of the range
// to: the ending value of the range
// targetNumTrailingZeros: the number of trailing zeros to check for
//
// This function should be called by the client as it assigns sub-jobs to the peers
func (p *Peer) InitJob(from int, to int, targetNumTrailingZeros int) {
	logger.Println("Job initiation request started by client with values:", from, to, targetNumTrailingZeros)

	// If too small, don't bother parallelizing
	if (to - from) < MAX_SIZE {
		p.StartJob(from, to, targetNumTrailingZeros)
		return
	}

	var wg sync.WaitGroup
	for start := from; start < to; start += MAX_SIZE {
		end := start + MAX_SIZE
		if end > to {
			end = to
		}

		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			// Assuming round-robin assignment of jobs to peers
			peerID := (s - from) / MAX_SIZE % len(p.Peers)
			p.Peers[peerID].StartJob(s, e, targetNumTrailingZeros)
		}(start, end)
	}

	wg.Wait()
}

// StartJob starts a job with the given parameters
func (p *Peer) StartJob(from int, to int, targetNumTrailingZeros int) {
	logger.Println("Job started by peer for the range", p.ID, from, to)
	p.StartCalculation(from, to, targetNumTrailingZeros)
}

// TODO: Update this function to calculate the proof of work instead. I added a dummy implementation for now to test the parallelization
//
//	Also, if we want other peers to stop once a peer finds a value, we can use a channel to communicate between peers to stop
func (p *Peer) StartCalculation(from int, to int, targetNumTrailingZeros int) {
	newBlock := Block{
		Nonce: from,
	}

	for !newBlock.IsValidBlock(targetNumTrailingZeros) {
		if newBlock.Nonce > to {
			break
		}
		newBlock.Nonce++
		logger.Printf("Trying nonce %d\n", newBlock.Nonce)
		newBlock.Hash = CalculateHash(newBlock.Nonce)
	}

	if newBlock.IsValidBlock(targetNumTrailingZeros) {
		logger.Printf("Found a valid block: Nonce %d, Hash %s, Peer ID %d\n", newBlock.Nonce, newBlock.Hash, p.ID)
	} else {
		logger.Printf("No valid block found within the range %d to %d \n", from, to)
	}

}

func (b Block) IsValidBlock(targetNumTrailingZeros int) bool {
	return strings.HasPrefix(b.Hash, strings.Repeat("0", targetNumTrailingZeros))
}

func CalculateHash(nonce int) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d", nonce)))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func setupBenchmark() {
	logger.SetOutput(io.Discard) // Discard log output during benchmarks
}
