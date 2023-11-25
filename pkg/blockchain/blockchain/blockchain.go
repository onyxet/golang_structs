package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	blockchain "structs/pkg/blockchain/db"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
	db           blockchain.Database
}

func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	newBlockChain := Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
		blockchain.Database{},
	}
	newBlockChain.db.CreateDB()
	newBlockChain.db.Init()
	return newBlockChain
}

func (b *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.Mine(b.difficulty)
	b.chain = append(b.chain, newBlock)

	jsonData, err := json.Marshal(newBlock.data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = b.db.Connection.Exec("INSERT INTO blockchain (timestamp, data, prev_hash, hash) VALUES (?, ?, ?, ?)",
		newBlock.timestamp, string(jsonData), newBlock.previousHash, newBlock.hash)
	if err != nil {
		log.Fatal(err)
	}
}

func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
