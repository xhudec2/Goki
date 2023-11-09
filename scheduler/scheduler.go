package scheduler

import (
	"fmt"
	"log"
	. "project/database"
	"time"

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
const COLLAPSE_TIME = 1200

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

func (queues *Scheduler) FillScheduler(cards *Table[Card], today int) (IDsPtr *[]ID, err error) {
	IDs := make([]ID, 0, len(*cards))
	for key := range *cards {
		card := (*cards)[key]
		IDs = append(IDs, card.ID)
		switch (*cards)[key].Queue {
		case 0:
			queues.New.Push(&card)
		case 1, 3:
			if card.Due > int(time.Now().Unix())*1000+COLLAPSE_TIME {
				continue
			}
			queues.Learing.Push(&card)
		case 2:
			if card.Due > today {
				continue
			}
			queues.Repeat.Push(&card)
		case -1:
			fmt.Println("Suspended card: ", (*cards)[key].Queue)
			continue
		default:
			log.Fatal("incorrect card_q number of card: ", (*cards)[key].Queue)
		}
	}
	IDsPtr = &IDs
	return
}

func (queues *Scheduler) Study(cards *Table[Card], db *gorm.DB, flds *map[ID]StudyNote) (err error) {
	for i := 0; i < Q_SIZE; i++ {
		c, err := queues.GetCard(cards)
		if c == nil || err != nil {
			log.Fatal(err)
			return err
		}

		again, err := StudyCard(c, db, flds)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if again {
			fmt.Println("Card added to queue")
		}
	}
	return
}
