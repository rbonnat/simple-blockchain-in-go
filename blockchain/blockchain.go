package blockchain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain contains a slice of Block and a time function
type Blockchain struct {
	now    NowFunc
	Blocks []Block
}

var (
	// ErrInvalidBlock is returned when block is invalid
	ErrInvalidBlock = errors.New("invalid block")
)

// NowFunc defines a function that returns time.Time
type NowFunc func() time.Time

// New initializes and returns a Blockchain instance
func New(ctx context.Context, now NowFunc) *Blockchain {
	var b = Blockchain{
		now: now,
	}

	genesisBlock := Block{
		0,
		b.now().String(),
		0,
		"",
		"",
	}

	b.Blocks = append(b.Blocks, genesisBlock)

	return &b
}

// InsertNewBlock adds a block to the Blockchain after validating it
func (b *Blockchain) InsertNewBlock(ctx context.Context, value int) (*Block, error) {
	var err error

	lastBlock := b.Blocks[len(b.Blocks)-1]

	// Generate new block
	newBlock, err := b.generateBlock(lastBlock, value)
	if err != nil {
		return nil, err
	}

	// Validate new block
	err = validateBlock(*newBlock, lastBlock)
	if err != nil {
		return nil, err
	}

	// Add new block to Blockchain
	newBlockchain := append(b.Blocks, *newBlock)
	b.replaceChain(newBlockchain)

	return newBlock, nil
}

func (b *Blockchain) generateBlock(oldBlock Block, value int) (*Block, error) {
	var newBlock Block

	t := b.now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Value = value
	newBlock.PrevHash = oldBlock.Hash

	hash, err := generateHash(newBlock)
	if err != nil {
		return nil, err
	}
	newBlock.Hash = hash

	return &newBlock, nil
}

func (b *Blockchain) replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(b.Blocks) {
		b.Blocks = newBlocks
	}
}

// Print prints the Blockchain
func (b *Blockchain) Print() {
	spew.Dump(b.Blocks)
}

func generateHash(block Block) (string, error) {
	record := string(block.Index) + block.Timestamp + string(block.Value) + block.PrevHash
	h := sha256.New()

	_, err := h.Write([]byte(record))
	if err != nil {
		return "", err
	}

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed), nil
}

func validateBlock(newBlock, oldBlock Block) error {
	if oldBlock.Index+1 != newBlock.Index {
		return ErrInvalidBlock
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return ErrInvalidBlock
	}

	hash, err := generateHash(newBlock)
	if err != nil {
		return err
	}

	if hash != newBlock.Hash {
		return ErrInvalidBlock
	}

	return nil
}
