package blockchain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testTime = "2020-07-04T14:05:53-04:00"
)

func mockNowFunc(t *testing.T) func() time.Time {
	now, err := time.Parse(time.RFC3339, testTime)
	require.NoError(t, err)

	return func() time.Time { return now }
}

func TestNew(t *testing.T) {
	nowFunc := mockNowFunc(t)

	tests := map[string]struct {
		expected *Blockchain
	}{
		"Succeed and returns a pointer to blockchain": {
			&Blockchain{
				nowFunc,
				[]Block{
					{
						0,
						nowFunc().String(),
						0,
						"",
						"",
					},
				},
			},
		},
	}

	for name, test := range tests {
		b := New(context.TODO(), nowFunc)
		assert.IsType(t, test.expected, b, "Incorrect type returned in test: '%s'", name)
		assert.Equal(t, test.expected.Blocks, b.Blocks, "Blockchains are different: '%s'", name)
	}
}

func TestGenerateHash(t *testing.T) {
	nowFunc := mockNowFunc(t)
	tests := map[string]struct {
		block Block
		hash  string
		err   error
	}{
		"Succeed and return Hash": {
			Block{
				0,
				nowFunc().String(),
				0,
				"",
				"",
			},
			"ea49d2502e4f551c16b7e8bf3b21011f9d3ffa29586026cb822e476a121b0c13",
			nil,
		},
	}

	for name, test := range tests {
		b := test.block
		h, err := generateHash(b)

		assert.Equal(t, test.hash, h, "Hash is incorrect: '%s'", name)
		assert.Equal(t, test.err, err, "Error is incorrect: '%s'", name)
	}

}

func TestGenerateBlock(t *testing.T) {
	nowFunc := mockNowFunc(t)

	b := New(context.TODO(), nowFunc)
	tests := map[string]struct {
		oldBlock Block
		value    int

		newBlock Block
		err      error
	}{
		"Succeed and return newBlock": {
			Block{
				1,
				nowFunc().String(),
				5,
				"",
				"dca5519519c4ac3e3554f59517ccc4e755720a10bc13dca804150d4ce6370141",
			},
			10,
			Block{
				2,
				nowFunc().String(),
				10,
				"dca5519519c4ac3e3554f59517ccc4e755720a10bc13dca804150d4ce6370141",
				"872bd64d7d4bf0d900d33343bd5f2305d6fb4acadc1bf3e33c9a206f18cb1354",
			},
			nil,
		},
	}

	for name, test := range tests {
		oldBlock := test.oldBlock
		newBlock, err := b.generateBlock(oldBlock, test.value)

		assert.Equal(t, test.newBlock, *newBlock, "Blocks are not equal: '%s'", name)
		assert.Equal(t, test.err, err, "Error is invalid: '%s'", name)
	}

}

func TestReplaceChain(t *testing.T) {
	nowFunc := mockNowFunc(t)
	ctx := context.TODO()

	bc1 := New(ctx, nowFunc)
	bc1.InsertNewBlock(ctx, 4)

	bc2 := New(ctx, nowFunc)
	bc2.InsertNewBlock(ctx, 9)
	bc2.InsertNewBlock(ctx, 8)

	tests := map[string]struct {
		bc        *Blockchain
		newBlocks []Block
		expected  []Block
	}{
		"Succeed and return newBlock": {
			bc1,
			bc2.Blocks,
			bc2.Blocks,
		},
		"Succeed and return oldBlock": {
			bc2,
			bc1.Blocks,
			bc2.Blocks,
		},
	}

	for name, test := range tests {
		test.bc.replaceChain(test.newBlocks)
		assert.Equal(t, test.expected, test.bc.Blocks, "Wrong chain replacement: '%s'", name)
	}
}

func TestIsBlockValid_Fail(t *testing.T) {
	nowFunc := mockNowFunc(t)

	tests := map[string]struct {
		oldBlock Block
		newBlock Block

		err error
	}{
		"Failed, wrong Index": {
			Block{
				Index:     0,
				Timestamp: nowFunc().String(),
				Value:     0,
				Hash:      "",
				PrevHash:  "",
			},
			Block{
				Index:     2,
				Timestamp: nowFunc().String(),
				Value:     4,
				Hash:      "883ca4d8350168ffcacf7a84f655aed26c781a42806615c555f00f25fcaaf655",
				PrevHash:  "",
			},
			ErrInvalidBlock,
		},
		"Failed, wrong PrevHash": {
			Block{
				Index:     0,
				Timestamp: nowFunc().String(),
				Value:     0,
				Hash:      "",
				PrevHash:  "",
			},
			Block{
				Index:     1,
				Timestamp: nowFunc().String(),
				Value:     4,
				Hash:      "883ca4d8350168ffcacf7a84f655aed26c781a42806615c555f00f25fcaaf655",
				PrevHash:  "wrongprevioushash",
			},
			ErrInvalidBlock,
		},
		"Failed, wrong Hash": {
			Block{
				0,
				nowFunc().String(),
				0,
				"",
				"",
			},
			Block{
				1,
				nowFunc().String(),
				4,
				"wronghash",
				"",
			},
			ErrInvalidBlock,
		},
	}

	for name, test := range tests {
		err := validateBlock(test.newBlock, test.oldBlock)
		assert.Equal(t, test.err, err, "Block should be invalid: '%s'", name)
	}

}
