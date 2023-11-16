package window

import (
	"fmt"
	"log"
	"src/backend/database"
	"src/backend/scheduler"
	"src/backend/tables"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

func AddCard(app fyne.App) {
	w := app.NewWindow("Add Card")
	w.Resize(fyne.NewSize(600, 450))
	w.CenterOnScreen()
	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Add Card"),
		),
	)
	w.Show()
}

func CycleCards(data *AppData, index int) {
	id := (*data.Decks)[index].ID
	deck := (*data.Decks)[index]
	cards := make(database.Table[tables.Card], deck.New+deck.Due+deck.Learn)
	err := database.DBGetter[tables.Card](data.DB, &cards, fmt.Sprintf("did = %d", id))
	if err != nil {
		log.Fatal(err)
	}

	today := tables.TodayRelative(data.DB)
	scheduler := scheduler.InitScheduler()
	IDs, err := scheduler.FillScheduler(&cards, today)
	if err != nil {
		log.Fatal(err)
	}
	flds := make(map[tables.ID]tables.StudyNote, len(cards))
	err = tables.GetFlds(IDs, data.DB, &flds)
	if err != nil {
		return
	}

	studyFunc := func(card *tables.Card, db *gorm.DB, conf *tables.Config, flds *map[tables.ID]tables.StudyNote) {
		studyInterface(data, card, db, conf, flds, scheduler)
	}

	scheduler.Study(data.DB, &tables.CONFIG, &flds, studyFunc)
}

func studyInterface(data *AppData, card *tables.Card, db *gorm.DB, conf *tables.Config, flds *map[tables.ID]tables.StudyNote, scheduler *scheduler.Scheduler) {

	studyFunc := func(card *tables.Card, db *gorm.DB, conf *tables.Config, flds *map[tables.ID]tables.StudyNote) {
		studyInterface(data, card, db, conf, flds, scheduler)
	}

	clickFunc := func() {
		card.UpdateCard(tables.GOOD, data.DB, conf)
		scheduler.ScheduleCard(card, tables.TodayRelative(db))
		scheduler.Study(data.DB, conf, flds, studyFunc)

	}

	flip := func() {
		data.Window.SetContent(
			container.NewVBox(
				MenuButtons(data),
				layout.NewSpacer(),
				container.NewCenter(
					widget.NewLabel((*flds)[card.ID].Back),
				),
				layout.NewSpacer(),
				container.NewCenter(
					widget.NewButton("Next", clickFunc),
				),
			),
		)
	}

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewLabel((*flds)[card.ID].Front),
			),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewButton("Flip", flip),
			),
		),
	)
}
