package tables

import "gorm.io/gorm"

func (card *Card) updateNewCard(db *gorm.DB) {
	card.Queue = LEARNING
	card.Type = LEARNING
	card.Due = MINUTE
}

func (card *Card) updateLrnCard(db *gorm.DB) {
	card.Queue = LEARNING
	card.Type = LEARNING
	card.Due = MINUTE
}

func (card *Card) updateRevCard(db *gorm.DB) {
	card.Queue = LEARNING
	card.Type = LEARNING
	card.Due = MINUTE
}

func (card *Card) UpdateCard(db *gorm.DB, grade string) (addToQ bool) {
	card.Reps = 1
	if card.Left >= 1000 {
		addToQ = true
		card.Left -= 1000
	}
	if card.Left%10 != 0 {
		card.Left -= 1
	}

	db.Model(Card{}).Where("id = ?", card.ID).Updates(*card)
	return
}
