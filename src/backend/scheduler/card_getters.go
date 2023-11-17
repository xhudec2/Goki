package scheduler

import (
	"fmt"
	"log"
	. "src/backend/tables"

	q "github.com/daviddengcn/go-villa"
)

// The Pop method of PriorityQueue does not return nil if the queue is empty...
// This is a workaround
func pop(q *q.PriorityQueue) (*Card, error) {
	if q.Len() <= 0 {
		return nil, fmt.Errorf("no cards in queue")
	}
	cardItem := q.Pop()
	card, _ := cardItem.(*Card)
	return card, nil
}

// for now, it cycles cards sequentially,
// the order of Learing and Review could be randomized
func (queues *Scheduler) GetCard() (card *Card, err error) {

	card, err = pop(queues.New)
	if err == nil {
		return card, nil
	}

	card, err = pop(queues.Learing)
	if err == nil {
		return card, nil
	}

	card, err = pop(queues.Review)
	if err == nil {
		return card, nil
	}
	log.Println("no cards in queue")
	return nil, nil
}
