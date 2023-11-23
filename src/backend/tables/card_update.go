package tables

// Logic behind updates is the same as in the original Anki SRS algorithm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

const (
	AGAIN = 1
	HARD  = 2
	GOOD  = 3
	EASY  = 4
)

func (card *Card) updateNewCard(grade int, db *gorm.DB, config *Config) {
	card.Queue = LEARNING
	card.Type = LEARNING
	card.Factor = config.New.Factor
	switch grade {
	case AGAIN:
		card.Due = config.New.Delays[0] * MINUTE
	case HARD:
		card.Due = config.New.Delays[1] * MINUTE / 2
	case GOOD:
		card.Due = config.New.Delays[1] * MINUTE
	case EASY:
		card.graduateCard(1, db, config)
	default:
		log.Fatal("Invalid grade")
	}
	card.Left = len(config.New.Delays)*1000 + len(config.New.Delays)
}

func (card *Card) updateLrnCard(grade int, db *gorm.DB, config *Config) {
	index := len(config.New.Delays) - 1 - card.Left/1000
	index = min(len(config.New.Delays)-1, max(index-1, 0))
	fmt.Println(index)
	switch grade {
	case AGAIN:
		// go to the beginning
		card.Left = len(config.New.Delays) + card.Left%1000
		card.Due = MINUTE
	case HARD:
		// go back one step
		card.Left += 1000
		card.Due = config.New.Delays[index] * MINUTE
	case GOOD:
		card.Left -= 1001
		if card.Left%1000 <= 0 {
			card.graduateCard(0, db, config)
		} else {
			card.Due = config.New.Delays[index] * MINUTE
		}
	case EASY:
		card.graduateCard(1, db, config)
	default:
		log.Fatal("Invalid grade")
	}
}

func (card *Card) updateRevCard(grade int, db *gorm.DB, config *Config) {
	var factor float64
	switch grade {
	case AGAIN:
		card.revLapse(config)
		return
	case HARD:
		factor = 1 - config.Review.HardFactor
	case GOOD:
		factor = 0
	case EASY:
		factor = 1 - config.Review.EasyBonus
	default:
		log.Fatal("Invalid grade")
	}
	card.Factor += int(factor * 1000)
	card.Ivl = int(float64(card.Ivl) * float64(card.Factor) / 1000)
	card.Due = TodayRelative(db) + card.Ivl
}

func (card *Card) graduateCard(level int, db *gorm.DB, config *Config) {
	card.Queue = REVIEW
	card.Type = REVIEW
	card.Ivl = config.New.Ints[level]
	card.Due = TodayRelative(db) + card.Ivl
}

func (card *Card) revLapse(config *Config) {
	now := int(time.Now().Unix())

	card.Lapses++
	card.Ivl = int(float64(card.Ivl) * config.Lapse.Mult)
	card.Queue = LEARNING
	card.Type = LEARNING

	card.Due = now + config.Lapse.Delays[0]*MINUTE
}

func (card *Card) UpdateCard(grade int, db *gorm.DB, config *Config) {
	card.Reps += 1
	switch card.Queue {
	case NEW:
		card.updateNewCard(grade, db, config)
	case LEARNING, USER_SUSPENDED:
		card.updateLrnCard(grade, db, config)
	case REVIEW:
		card.updateRevCard(grade, db, config)
	}
	db.Model(Card{}).Where("id = ?", card.ID).Updates(*card)
}
