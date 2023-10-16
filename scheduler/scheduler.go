package scheduler

import (
	"database/sql"
	"project/database"

	q "github.com/Workiva/go-datastructures/queue"
)

type Scheduler struct {
	New     *q.Queue
	Learing *q.Queue
	Repeat  *q.Queue
}

const Q_SIZE = 32

func Scheduler_init() (qs *Scheduler) {
	qs.New = q.New(Q_SIZE)
	qs.Learing = q.New(Q_SIZE)
	qs.Repeat = q.New(Q_SIZE)
	return
}

func (queues *Scheduler) Fill_scheduler(db *sql.DB) (err error) {
	return
}

func Study_card(card database.Card, db *sql.DB) (err error) {
	return nil
}

func (queues *Scheduler) Study(db *sql.DB) (err error) {
	if len := queues.New.Len(); len > 0 {
		_ /* cards */, err := queues.New.Get(len)
		if err != nil {
			return err
		}
		//		cards = ([](database.Card))(cards)
		//		for _, card := range cards {
		//			Study_card(card, db)
		//		}
	}
	return
}
