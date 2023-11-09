package scheduler

import (
	"fmt"
	"log"
	. "project/database"

	q "github.com/daviddengcn/go-villa"
	"gorm.io/gorm"
)

type Scheduler struct {
	// these Q names are the same as those used in the .apkg database
	// will change them to suit my usage more in the future
	New     *q.PriorityQueue
	Learing *q.PriorityQueue
	Repeat  *q.PriorityQueue
}

const Q_SIZE = 32 // Q_SIZE == deck card limit ?
const DEFAULT_WHERE = ""

func InitScheduler() (qs *Scheduler) {
	cmp := func(a, b interface{}) int {
		card := a.(*Card)
		other := b.(*Card)
		if card.Due < other.Due {
			return 1
		} else if card.Due > other.Due {
			return -1
		}
		return 0
	}
	return &Scheduler{
		New:     q.NewPriorityQueueCap(cmp, Q_SIZE),
		Learing: q.NewPriorityQueueCap(cmp, Q_SIZE),
		Repeat:  q.NewPriorityQueueCap(cmp, Q_SIZE),
	}
}

func (queues *Scheduler) FillScheduler(cards *Table[Card]) (IDsPtr *[]ID, err error) {
	IDs := make([]ID, 0, len(*cards))
	for key := range *cards {
		card := (*cards)[key]
		switch (*cards)[key].Queue {
		case 0:
			queues.New.Push(&card)
		case 1, 3:
			queues.Learing.Push(&card)
		case 2:
			queues.Repeat.Push(&card)
		case -1:
			continue
		default:
			_, err = fmt.Println("Incorrect card_q number of card: ", (*cards)[key].Queue)
			log.Fatal(err)
			return nil, err
		}
		IDs = append(IDs, card.ID)
	}
	IDsPtr = &IDs
	return
}

func (queues *Scheduler) Study(cards *Table[Card], db *gorm.DB, flds *map[ID]StudyNote) (err error) {
	for i := 0; i < Q_SIZE; i++ {
		c, err := queues.getCard(cards)
		if c == nil || err != nil {
			log.Fatal(err)
			return err
		}

		if err = StudyCard(c, db, flds); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return
}

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

func (queues *Scheduler) getCard(cards *Table[Card]) (card *Card, err error) {

	card, err = pop(queues.New)
	if err == nil {
		return card, nil
	}

	card, err = pop(queues.Learing)
	if err == nil {
		return card, nil
	}

	card, err = pop(queues.Repeat)
	if err == nil {
		return card, nil
	}
	return nil, fmt.Errorf("no cards in queues")
}
