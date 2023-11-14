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

const Q_SIZE = 64
const DEFAULT_WHERE = ""
const COLLAPSE_TIME = 1200 // the time a card can be studied in advance

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
	case NEW:
		queues.New.Push(card)
	case LEARNING, USER_SUSPENDED:
		if card.Due > int(time.Now().Unix())*1000+COLLAPSE_TIME {
			return false
		}
		queues.Learing.Push(card)
	case REVIEW:
		if card.Due > today {
			return false
		}
		queues.Review.Push(card)
	case SUSPENDED:
		fmt.Println("Suspended card: ", card.ID)
	default:
		log.Fatal("invalid card queue: ", card.Queue)
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
		err = StudyCard(card, db, conf, flds)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if queues.ScheduleCard(card, TodayRelative(db)) {
			// only a debug print
			fmt.Println("Card added to queue")
		}
	}
	return
}
