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

	askQueue.Add(buildOrder(1, "ask", 0.041))

	t.Logf("askQueue.Len() = %d", askQueue.Len())
	t.Logf("askQueue.Size() = %d", askQueue.Size())

	askQueue.Add(buildOrder(2, "ask", 0.045))
	askQueue.Add(buildOrder(3, "ask", 0.047))
	askQueue.Add(buildOrder(8, "ask", 0.045))
	askQueue.Add(buildOrder(10, "ask", 0.049))
	askQueue.Add(buildOrder(20, "ask", 0.045))
	askQueue.Add(buildOrder(21, "ask", 0.041))

	t.Logf("askQueue.Len() = %d", askQueue.Len())
	t.Logf("askQueue.Size() = %d", askQueue.Size())

	for order := range askQueue.All() {
		t.Logf("order = %+v", order)
	}

	apr := decimal.NewFromFloat(0.045)
	f := func(order Order) bool {
		return order.Apr.GreaterThan(apr)
	}

	for order := range askQueue.IterWith(f) {
		t.Logf("filtered order = %+v", order)
	}
}

func buildOrder(id uint64, side string, apr float64) Order {
	return Order{
		ID:        id,
		Side:      side,
		Asset:     "Ethereum_ETH",
		Term:      "30D",
		Amount:    1,
		Apr:       decimal.NewFromFloat(apr),
		Timestamp: uint64(time.Now().Unix()),
	}
}

func TestBidQueue_Add(t *testing.T) {
	bidQueue := NewBidQueue()

	t.Logf("bidQueue.Len() = %d", bidQueue.Len())
	t.Logf("bidQueue.Size() = %d", bidQueue.Size())

	bidQueue.Add(buildOrder(1, "bid", 0.041))

	t.Logf("bidQueue.Len() = %d", bidQueue.Len())
	t.Logf("bidQueue.Size() = %d", bidQueue.Size())

	bidQueue.Add(buildOrder(2, "bid", 0.045))
	bidQueue.Add(buildOrder(3, "bid", 0.047))
	bidQueue.Add(buildOrder(10, "bid", 0.049))

	t.Logf("bidQueue.Len() = %d", bidQueue.Len())
	t.Logf("bidQueue.Size() = %d", bidQueue.Size())

	bidQueue.Remove(buildOrder(3, "bid", 0.041))
	bidQueue.Remove(buildOrder(11, "bid", 0.041))

	t.Logf("bidQueue.Len() = %d", bidQueue.Len())
	t.Logf("bidQueue.Size() = %d", bidQueue.Size())

	for order := range bidQueue.All() {
		t.Logf("order = %+v", order)
	}

	apr := decimal.NewFromFloat(0.045)
	f := func(order Order) bool {
		return order.Apr.LessThanOrEqual(apr)
	}

	for order := range bidQueue.IterWith(f) {
		t.Logf("filtered order = %+v", order)
	}
}
