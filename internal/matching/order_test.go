package matching

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestAskQueue(t *testing.T) {
	askQueue := NewAskQueue()

	t.Logf("askQueue.Len() = %d", askQueue.Len())
	t.Logf("askQueue.Size() = %d", askQueue.Size())

	askQueue.Add(buildOrder(1, 0.041))

	t.Logf("askQueue.Len() = %d", askQueue.Len())
	t.Logf("askQueue.Size() = %d", askQueue.Size())

	askQueue.Add(buildOrder(2, 0.045))
	askQueue.Add(buildOrder(3, 0.047))
	askQueue.Add(buildOrder(10, 0.049))

	t.Logf("askQueue.Len() = %d", askQueue.Len())
	t.Logf("askQueue.Size() = %d", askQueue.Size())
}

func buildOrder(id uint64, apr float64) Order {
	return Order{
		ID:        id,
		Side:      "ask",
		Asset:     "Ethereum_ETH",
		Term:      "30D",
		Amount:    1,
		Apr:       decimal.NewFromFloat(apr),
		Timestamp: uint64(time.Now().Unix()),
	}
}
