package scheduler

import (
	"fmt"
	"log"
	. "project/database"

	q "github.com/Workiva/go-datastructures/queue"
	"gorm.io/gorm"
)

type Scheduler struct {
	// these Q names are the same as those used in the .apkg database
	// will change them to suit my usage more in the future
	New     *q.Queue
	Learing *q.Queue
	Repeat  *q.Queue
}

const Q_SIZE = 32 // Q_SIZE == deck card limit ?
const DEFAULT_WHERE = ""

func SchedulerInit() (qs *Scheduler) {
	return &Scheduler{
		New:     q.New(Q_SIZE),
		Learing: q.New(Q_SIZE),
		Repeat:  q.New(Q_SIZE),
	}
}

func (queues *Scheduler) FillScheduler(cards *Table[Card]) (err error) {
	for key := range *cards {
		switch (*cards)[key].Queue {
		case 0:
			queues.New.Put((*cards)[key].ID)
		case 1, 3:
			queues.Learing.Put((*cards)[key].ID)
		case 2:
			queues.Repeat.Put((*cards)[key].ID)
		default:
			fmt.Println("Incorrect card_q number of card: ", (*cards)[key].Queue)
			return
		}
	}
	return
}

func StudyQ(q *q.Queue, cards *Table[Card], db *gorm.DB) (err error) {
	len := q.Len()
	if len <= 0 {
		return
	}
	qItems, err := q.Get(len)
	if err != nil {
		return err
	}
	IDs := make([]ID, len)
	for i, id := range qItems {
		id, ok := id.(ID)
		IDs[i] = id
		if !ok {
			fmt.Println("card nok: ", id)
			continue
		}
	}
	flds := make(map[ID]StudyNote, len)
	GetFlds(IDs, db, &flds)

	for _, id := range IDs {
		card := (*cards)[id]
		err = StudyCard(&card, db, &flds)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}

	return
}

func (queues *Scheduler) Study(cards *Table[Card], db *gorm.DB) (err error) {
	fmt.Println("New")
	err = StudyQ(queues.New, cards, db)
	if err != nil {
		fmt.Println("New failed")
	}
	fmt.Println("Learning")
	err = StudyQ(queues.Learing, cards, db)
	if err != nil {
		fmt.Println("Learning failed")
	}
	fmt.Println("Repeat")
	err = StudyQ(queues.Repeat, cards, db)
	if err != nil {
		fmt.Println("Repeat failed")
	}
	return
}
