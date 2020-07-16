package service

import (
	"context"

	"github.com/rbonnat/blockchain-in-go/blockchain"
)

// Service contains data for services
type Service struct {
	blockchain *blockchain.Blockchain
}

// New initializes and returns a pointer to an instance of service
func New(ctx context.Context, now blockchain.NowFunc) *Service {
	var s Service
	s.blockchain = blockchain.New(ctx, now)

	return &s
}

// Blocks is a service that returns the blocks of the Blockchain
func (s *Service) Blocks(ctx context.Context) []blockchain.Block {
	return s.blockchain.Blocks
}

// InsertNewBlock inserts a new block into the Blockchain
func (s *Service) InsertNewBlock(ctx context.Context, BPM int) (*blockchain.Block, error) {
	return s.blockchain.InsertNewBlock(ctx, BPM)
}
