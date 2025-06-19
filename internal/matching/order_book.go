package matching

import (
	"container/list"

	"github.com/shopspring/decimal"
)

type OrderBook struct {
	supplys AskQueue
	borrows BidQueue
}

type Order struct {
	// Order ID
	ID        uint64
	Side      string
	Asset     string
	Term      string
	Amount    float64
	Apr       decimal.Decimal
	Timestamp uint64
}

func NewAskQueue() *AskQueue {
	return &AskQueue{
		rootList: list.New(),
		aprMap:   make(map[decimal.Decimal]*list.Element),
	}
}

// AskQueue is a priority queue that sorts supply orders in-memory.
type AskQueue struct {
	// RootList is a two-dimensional list that sorts orders
	// by APR in ascending order and timestamp in ascending order.
	rootList *list.List

	// Map of APR to list.Element.
	aprMap map[decimal.Decimal]*list.Element

	// The total number of orders in the queue.
	totalOrders int
}

func (q *AskQueue) Add(order Order) {
	apr := order.Apr

	// If the root list is empty, create a new list and add the order.
	if q.rootList.Len() == 0 {
		aprList := list.New()
		aprList.PushBack(order)

		element := q.rootList.PushBack(aprList)
		q.aprMap[apr] = element

		q.totalOrders++

		return
	}

	aprElement, ok := q.aprMap[apr]
	// If the APR is already in the map, add the order to the list.
	if ok {
		aprList := aprElement.Value.(*list.List)
		aprList.PushBack(order)
	} else {
		// If the APR is not in the map, create a new list and add the order.
		newAprList := list.New()
		newAprList.PushBack(order)

		// Find the correct position to insert the root list.
		for e := q.rootList.Front(); ; e = e.Next() {
			if e == nil {
				// Place the new APR list at the end of the root list.
				element := q.rootList.PushBack(newAprList)
				q.aprMap[apr] = element

				break
			}

			aprListInRoot := e.Value.(*list.List)

			// Find the first APR in the front of the sub-list that is greater than the new APR.
			if aprListInRoot.Front().Value.(Order).Apr.GreaterThan(apr) {
				if e.Prev() == nil {
					// Place the new APR list at the front of the root list.
					q.rootList.PushFront(newAprList)
				} else {
					// Place the new APR list before the current APR list.
					newAprPosition := q.rootList.InsertBefore(newAprList, e)
					q.aprMap[apr] = newAprPosition
				}

				break
			}
		}
	}

	q.totalOrders++
}

// Len returns the inner root list length.
func (q *AskQueue) Len() int {
	return q.rootList.Len()
}

// Size returns the total number of orders in the queue.
func (q *AskQueue) Size() int {
	return q.totalOrders
}

// bid queue is a priority queue that sorts orders
// by time in ascending order
type BidQueue struct {
}
