package scheduler

import (
	"database/sql"
	"fmt"
	"project/database"
	"strings"

	q "github.com/Workiva/go-datastructures/queue"
)

type Scheduler struct {
	// these Q names are the same as those used in the .apkg database
	// will change them to suit my usage more in the future
	New     *q.Queue
	Learing *q.Queue
	Repeat  *q.Queue
}

const Q_SIZE = 32 // Q_SIZE == deck card limit ?

func Scheduler_init() (qs *Scheduler) {
	return &Scheduler{
		New:     q.New(Q_SIZE),
		Learing: q.New(Q_SIZE),
		Repeat:  q.New(Q_SIZE),
	}
}

func (queues *Scheduler) Fill_scheduler(card_ids database.Card_ids, db *sql.DB) (err error) {
	rows, err := db.Query(fmt.Sprintf("SELECT id, queue FROM cards WHERE id IN (%s)", strings.Join(card_ids, ",")))
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	for rows.Next() {
		var id, card_q string
		err = rows.Scan(&id, &card_q)
		if err != nil {
			fmt.Println("Error scanning rows")
			return
		}
		switch card_q {
		case "0":
			queues.New.Put(id)
		case "1", "3":
			queues.Learing.Put(id)
		case "2":
			queues.Repeat.Put(id)
		default:
			fmt.Println("Incorrect card_q number of card: ", card_q)
			return
		}
	}
	return
}

func Study_card(card string, cards database.Cards, db *sql.DB) (err error) {
	// TODO: unfinished, only prints the card for now
	// It should also change the db based on the "grade" given to the card while studying
	// return the grade as well to know where to store the card next
	fmt.Println(cards[database.Id(card)])
	return nil
}

func Study_q(q *q.Queue, cards database.Cards, db *sql.DB) (err error) {
	len := q.Len()
	if len <= 0 {
		return
	}

	q_items, err := q.Get(len)
	if err != nil {
		return err
	}

	for _, card := range q_items {
		card, ok := card.(string)
		if !ok {
			fmt.Println("card nok: ", card)
			continue
		}
		Study_card(card, cards, db)
	}

	return
}

func (queues *Scheduler) Study(cards database.Cards, db *sql.DB) (err error) {
	fmt.Println("New\n")
	err = Study_q(queues.New, cards, db)
	if err != nil {
		fmt.Println("New failed")
	}
	fmt.Println("Learning\n")
	err = Study_q(queues.Learing, cards, db)
	if err != nil {
		fmt.Println("Learning failed")
	}
	fmt.Println("Repeat\n")
	err = Study_q(queues.Repeat, cards, db)
	if err != nil {
		fmt.Println("Repeat failed")
	}
	return
}
