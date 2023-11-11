package scheduler

import (
	"fmt"
	"log"
	. "project/database"
	. "project/tables"
	"time"

	q "github.com/daviddengcn/go-villa"
	"gorm.io/gorm"
)

type Scheduler struct {
	New     *q.PriorityQueue
	Learing *q.PriorityQueue
	Review  *q.PriorityQueue
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
		Review:  q.NewPriorityQueueCap(cmp, Q_SIZE),
	}
}

func (queues *Scheduler) ScheduleCard(card *Card, today int) bool {
	switch card.Queue {
	case 0:
		queues.New.Push(card)
	case 1, 3:
		if card.Due > int(time.Now().Unix())*1000+COLLAPSE_TIME {
			return false
		}
		queues.Learing.Push(card)
	case 2:
		if card.Due > today {
			return false
		}
		queues.Review.Push(card)
	case -1:
		fmt.Println("Suspended card: ", card.Queue)
	default:
		log.Fatal("incorrect card_q: ", card.Queue)
	}
	return true
}

func (queues *Scheduler) FillScheduler(cards *Table[Card], today int) (IDsPtr *[]ID, err error) {
	IDs := make([]ID, 0, len(*cards))
	for key := range *cards {
		card := (*cards)[key]
		if queues.ScheduleCard(&card, today) {
			IDs = append(IDs, card.ID)
		}
	}
	IDsPtr = &IDs
	return
}

func (queues *Scheduler) Study(cards *Table[Card], db *gorm.DB, conf *Config, flds *map[ID]StudyNote) (err error) {
	for i := 0; i < Q_SIZE; i++ {
		card, err := queues.GetCard(cards)
		if card == nil || err != nil {
			log.Fatal(err)
			return err
		}
		again, err := StudyCard(card, db, conf, flds)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if again {
			queues.ScheduleCard(card, TodayRelative(db))
			fmt.Println("Card added to queue")
		}
	}
	return
}
