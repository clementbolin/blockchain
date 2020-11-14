package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

// Block is one block of Blockchain
type Block struct {
	index 		int
	timeStamp 	string
	bpm 		int
	hash 		string
	prevHash	string
}

// Create hash
func calculHash(block Block) string {
	record := string(block.index) + block.timeStamp + string(block.bpm) + block.prevHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// createBlock create new Block structure
func createBlock(lastBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.index = lastBlock.index + 1
	newBlock.timeStamp = t.String()
	newBlock.bpm = BPM
	newBlock.prevHash = lastBlock.hash
	newBlock.hash = calculHash(newBlock)
	return newBlock, nil
}

// isBlockValid check if newBlock is valid
func isBlockValid(newBlock, lastBlock Block) error {
	if lastBlock.index + 1 != newBlock.index {
		return errors.New("Invalid index")
	}
	if lastBlock.hash != newBlock.prevHash {
		return errors.New("Invalid prev Hash")
	}
	if calculHash(newBlock) != newBlock.hash {
		return errors.New("Invalid Hash")
	}
	return nil
}

func main() {
	return
}
