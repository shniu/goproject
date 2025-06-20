package matching

import (
	"container/list"
	"iter"

	"github.com/shopspring/decimal"
)

type OrderBook struct {
	supplys AskQueue
	borrows BidQueue
}

type Order struct {
	// SeqId uint64 ?
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

func (q *AskQueue) All() iter.Seq[Order] {
	return func(yield func(Order) bool) {
		for e := q.rootList.Front(); e != nil; e = e.Next() {
			aprList := e.Value.(*list.List)
			for e := aprList.Front(); e != nil; e = e.Next() {
				order := e.Value.(Order)

				if !yield(order) {
					return
				}
			}
		}
	}
}

// IterWith returns iterator that stops when condition is met
func (q *AskQueue) IterWith(stop func(Order) bool) iter.Seq[Order] {
	return func(yield func(Order) bool) {
		for e := q.rootList.Front(); e != nil; e = e.Next() {
			aprList := e.Value.(*list.List)
			for orderE := aprList.Front(); orderE != nil; orderE = orderE.Next() {
				order := orderE.Value.(Order)
				if stop(order) {
					return
				}
				if !yield(order) {
					return
				}
			}
		}
	}
}

func NewBidQueue() *BidQueue {
	return &BidQueue{
		rootList:    list.New(),
		orderMap:    make(map[uint64]*list.Element),
		totalOrders: 0,
	}
}

// bid queue is a priority queue that sorts orders
// by time in ascending order
type BidQueue struct {
	rootList    *list.List
	orderMap    map[uint64]*list.Element
	totalOrders uint64
}

// Add adds an order into the queue.
func (q *BidQueue) Add(order Order) error {
	e := q.rootList.PushBack(order)
	q.orderMap[order.ID] = e
	q.totalOrders++
	return nil
}

// Remove removes an order from the queue.
func (q *BidQueue) Remove(order Order) error {
	e, ok := q.orderMap[order.ID]
	if ok {
		q.rootList.Remove(e)
		delete(q.orderMap, order.ID)
		q.totalOrders--
	}
	return nil
}

func (q *BidQueue) All() iter.Seq[Order] {
	return func(yield func(Order) bool) {
		for e := q.rootList.Front(); e != nil; e = e.Next() {
			order := e.Value.(Order)
			if !yield(order) {
				return
			}
		}
	}
}

func (q *BidQueue) IterWith(filter func(Order) bool) iter.Seq[Order] {
	return func(yield func(Order) bool) {
		for e := q.rootList.Front(); e != nil; e = e.Next() {
			order := e.Value.(Order)
			if filter(order) {
				if !yield(order) {
					return
				}
			}
		}
	}
}

func (q *BidQueue) Len() int {
	return q.rootList.Len()
}

func (q *BidQueue) Size() uint64 {
	return q.totalOrders
}
