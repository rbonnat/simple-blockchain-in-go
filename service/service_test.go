package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rbonnat/blockchain-in-go/blockchain"
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
	ctx := context.TODO()

	tests := map[string]struct {
		expected *Service
	}{
		"Succeed and returns a pointer to a Service": {
			&Service{
				blockchain.New(ctx, nowFunc),
			},
		},
	}

	for name, test := range tests {
		s := New(ctx, nowFunc)
		assert.IsType(t, test.expected, s, "Type is different: '%s'", name)
		assert.Equal(t, test.expected.Blocks(ctx), s.Blocks(ctx), "Blockchain are different: '%s'", name)
	}
}

func TestInsertNewBlock(t *testing.T) {
	nowFunc := mockNowFunc(t)
	ctx := context.TODO()

	tests := map[string]struct {
		expected *Service
	}{
		"Succeed and returns a pointer to a Service": {
			&Service{
				blockchain.New(ctx, nowFunc),
			},
		},
	}

	for name, test := range tests {
		s := New(ctx, nowFunc)
		assert.IsType(t, test.expected, s, "Type is different: '%s'", name)
		assert.Equal(t, test.expected.Blocks(ctx), s.Blocks(ctx), "Blockchain are different: '%s'", name)
	}
}
