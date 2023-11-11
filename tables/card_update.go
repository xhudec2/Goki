package tables

import (
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
	card.Left = 2000 + len(config.New.Delays)
}

func (card *Card) updateLrnCard(grade int, db *gorm.DB, config *Config) (reschedule bool) {
	reschedule = true
	switch grade {
	case AGAIN:
		// card is reset and needed to be studied from the beginning
		card.Left = 2000 + card.Left%1000
		card.Due = MINUTE
	case HARD:
		card.Due = config.Lapse.Delays[0] * MINUTE
		// add more
	case GOOD:
		card.Left--
		if card.Left%1000 <= 0 {
			card.graduateCard(0, db, config)
			reschedule = false
		} else {
			card.Due = config.New.Delays[card.Left%1000-1] * MINUTE
		}
	case EASY:
		card.graduateCard(1, db, config)
		reschedule = false
	default:
		log.Fatal("Invalid grade")
	}
	return
}

func (card *Card) updateRevCard(grade int, db *gorm.DB, config *Config) {
	var factor float64
	switch grade {
	case AGAIN:
		card.revLapse(config)
		return
	case HARD:
		factor = config.Review.HardFactor
	case GOOD:
		factor = config.Review.EasyFactor
	case EASY:
		factor = config.Review.EasyFactor + config.Review.EasyBonus
	default:
		log.Fatal("Invalid grade")
	}
	card.Ivl = int(float64(card.Ivl) * factor)
	card.Due = TodayRelative(db) + card.Ivl/1000
}

func (card *Card) graduateCard(level int, db *gorm.DB, config *Config) {
	card.Queue = REVIEW
	card.Type = REVIEW
	card.Factor = config.New.Factor - 200*card.Lapses
	card.Due = TodayRelative(db) + config.New.Ints[level]
	card.Ivl = 1
}

func (card *Card) revLapse(config *Config) {
	now := int(time.Now().Unix())

	card.Lapses++
	card.Ivl = 0
	card.Queue = LEARNING
	card.Type = LEARNING
	card.Due = now + config.Lapse.Delays[0]*MINUTE
}

func (card *Card) UpdateCard(grade int, db *gorm.DB, config *Config) (addToQ bool) {
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
	return
}
