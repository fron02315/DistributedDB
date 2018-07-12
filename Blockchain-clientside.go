package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
	"fmt"
)

// Block keeps block headers
type Block struct {
  Index int
	Timestamp     int64
  Nodeid1    string
  Nodeid2    string
  MsgSize       string
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  index := []byte(strconv.Itoa(b.Index))
	headers := bytes.Join([][]byte{b.PrevBlockHash, index, []byte(b.Nodeid1), []byte(b.Nodeid2), []byte(b.MsgSize), timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func isBlockValid(newBlock Block, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if bytes.Compare(oldBlock.Hash, newBlock.PrevBlockHash) != 0{
		return false
	}

	return true
}

// NewBlock creates and returns Block
func NewBlock(Nodeid1 string, Nodeid2 string, MsgSize string, oldBlock *Block) *Block {
  newIndex := oldBlock.Index+1
	block := &Block{newIndex, time.Now().Unix(), Nodeid1, Nodeid2, MsgSize, oldBlock.Hash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
  //To begin the new block, we assign the index=0 and all elements are null or zero
  genesisBlock := &Block{0, time.Now().Unix(), "0", "0", "0", []byte(""), []byte{}}
  genesisBlock.SetHash()
	return genesisBlock
}

type Blockchain struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func AddBlock(bc *Blockchain, Nodeid1 string, Nodeid2 string, MsgSize string) {
  //old block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(Nodeid1, Nodeid2, MsgSize, prevBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func PrintBC(bc *Blockchain){
  for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
    fmt.Printf("Block index: %d\n", block.Index)
		fmt.Printf("Node ID1: %s\n", block.Nodeid1)
    fmt.Printf("Node ID2: %s\n", block.Nodeid2)
    fmt.Printf("Message size: %s\n", block.MsgSize)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

/*func main() {


	bc.AddBlock("p141256", "8866ab3", 24)
  PrintBC(bc)

} */
