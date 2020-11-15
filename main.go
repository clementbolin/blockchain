package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

/*                  BlockChain                              */

// Block is one block of Blockchain
type Block struct {
	index 		int
	timeStamp 	string
	bpm 		int
	hash 		string
	prevHash	string
}

// Message contain body request value
type Message struct {
	BPM int
}

// BlockChain BlockChain
var BlockChain []Block
var mutex = &sync.Mutex{}

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
func isBlockValid(newBlock, lastBlock Block) bool {
	if lastBlock.index + 1 != newBlock.index {
		return false
	}
	if lastBlock.hash != newBlock.prevHash {
		return false
	}
	if calculHash(newBlock) != newBlock.hash {
		return false
	}
	return true
}

// replaceChain replace BlockChain by newBlockChain
func replaceChain(newBlock []Block) {
	if len(newBlock) > len(BlockChain) {
		BlockChain = newBlock
	}
}

/*                  BlockChain                              */

/*                  Web Server                              */

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// handleGetBlockchain get blockChain
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(BlockChain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// handleWriteBlock create block in BlockChain
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := createBlock(BlockChain[len(BlockChain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, BlockChain[len(BlockChain)-1]) {
		newBlockChain := append(BlockChain, newBlock)
		replaceChain(newBlockChain)
		spew.Dump(BlockChain)
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

// createRouter create router
func createRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

// run run web server
func run() {
	mux := createRouter()
	httpAddr := os.Getenv("PORT")
	log.Printf("HTTP Server Listening on port : %s !", httpAddr)
	s := &http.Server{
		Addr: ":" + httpAddr,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

/*                  Web Server                              */

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), 0, calculHash(genesisBlock), ""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		BlockChain = append(BlockChain, genesisBlock)
		mutex.Unlock()
	}()
	run()
}
