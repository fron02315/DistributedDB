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
	Timestamp     string
  Nodeid1    string
  Nodeid2    string
  Operation   string
  MsgSize       string
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
  index := []byte(strconv.Itoa(b.Index))
	headers := bytes.Join([][]byte{b.PrevBlockHash, index, []byte(b.Operation), []byte(b.Nodeid1), []byte(b.Nodeid2), []byte(b.MsgSize), []byte(b.Timestamp)}, []byte{})
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
func NewBlock(Nodeid1 string, Nodeid2 string, Operation string, MsgSize string, oldBlock *Block) *Block {
  newIndex := oldBlock.Index+1
	t := time.Now()
	block := &Block{newIndex, t.String(), Nodeid1, Nodeid2, Operation, MsgSize, oldBlock.Hash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	t := time.Now()
  //To begin the new block, we assign the index=0 and all elements are null or zero
  genesisBlock := &Block{0, t.String(), "0", "0", "null", "0", []byte(""), []byte{}}
  genesisBlock.SetHash()
	return genesisBlock
}

type Blockchain struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func AddBlock(bc *Blockchain, Nodeid1 string, Nodeid2 string, Operation string, MsgSize string) {
  //old block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(Nodeid1, Nodeid2, Operation, MsgSize, prevBlock)
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
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Node ID1: %s\n", block.Nodeid1)
    fmt.Printf("Node ID2: %s\n", block.Nodeid2)
    fmt.Printf("Operation performed: %s\n", block.Operation)
    fmt.Printf("Message size (Bytes): %s\n", block.MsgSize)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

/*func main() {


	bc.AddBlock("p141256", "8866ab3", 24)
  PrintBC(bc)

} */
