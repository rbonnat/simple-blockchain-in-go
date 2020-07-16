package blockchain

// Block describes describes a block of the Blockchain
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
	PrevHash  string `json:"prev_hash"`
	Hash      string `json:"hash"`
}
