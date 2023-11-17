package window

import (
	"fmt"
	"image/color"
	"log"
	"src/backend/database"
	"src/backend/scheduler"
	"src/backend/tables"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AddCard(data *Data) {
	w := data.App.NewWindow("Add Card")
	w.Resize(fyne.NewSize(600, 450))
	w.CenterOnScreen()

	frontInput := widget.NewEntry()
	BackInput := widget.NewEntry()
	frontInput.SetPlaceHolder("Enter text...")
	BackInput.SetPlaceHolder("Enter text...")

	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(200, 50))

	deck := make(chan string, 1)

	deckSelect := container.NewVBox(
		widget.NewLabel("Choose Deck"),
		widget.NewSelect(
			tables.DeckNames(data.StudyData.Decks),
			func(s string) {
				deck <- s
			},
		),
	)

	front := container.NewStack(
		bg,
		container.NewGridWithRows(
			2,
			widget.NewLabel("Front"),
			frontInput,
		),
	)
	back := container.NewStack(
		bg,
		container.NewGridWithRows(
			2,
			widget.NewLabel("Back"),
			BackInput,
		),
	)

	w.SetContent(
		container.NewVBox(
			layout.NewSpacer(),
			deckSelect,
			layout.NewSpacer(),
			container.NewHBox(
				layout.NewSpacer(),
				front,
				back,
				layout.NewSpacer(),
			),
			layout.NewSpacer(),
			widget.NewButton("Add",
				func() {
					did := tables.GetDeckID(data.StudyData.Decks, <-deck)
					database.AddNewCard(did, frontInput.Text, BackInput.Text, data.StudyData.DB)
					w.Close()
					Draw(data)
				}),
			layout.NewSpacer(),
		),
	)
	w.Show()
}

func CycleDeck(data *Data, deck *tables.Deck) {
	initDeckData(data, deck)

	studyFunc := func(card *tables.Card, studyData *scheduler.StudyData) {
		if card == nil {
			data.Window.SetContent(NothingToStudy(data))
			return
		}
		studyCard(data, card)
	}
	data.StudyData.Conf = &deck.Conf
	data.StudyData.StudyFunc = studyFunc
	data.StudyData.Scheduler.Study(data.StudyData)
}

func initDeckData(data *Data, deck *tables.Deck) {
	db := data.StudyData.DB

	id := deck.ID
	cards := database.Table[tables.Card]{}
	err := database.DBGetter[tables.Card](db, &cards, "", fmt.Sprintf("did = %d", id))

	if err != nil {
		log.Fatal(err)
	}

	queues := scheduler.InitScheduler()
	IDs, err := queues.FillScheduler(&cards, tables.TodayRelative(db))
	if err != nil {
		log.Fatal(err)
	}
	data.StudyData.Scheduler = queues
	flds := make(map[tables.ID]tables.StudyNote, len(cards))
	err = tables.GetFlds(IDs, db, &flds)
	if err != nil {
		log.Fatal(err)
	}
	data.StudyData.Flds = &flds
}

func studyCard(data *Data, card *tables.Card) {

	cardFront := (*data.StudyData.Flds)[card.ID].Front

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewLabel(cardFront),
			),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewButton("Flip", func() { flipClick(data, card) }),
			),
		),
	)
}

func gradeClick(data *Data, card *tables.Card, grade int) {
	card.UpdateCard(grade, data.StudyData.DB, data.StudyData.Conf)
	data.StudyData.Scheduler.ScheduleCard(card, tables.TodayRelative(data.StudyData.DB))
	data.StudyData.Scheduler.Study(data.StudyData)
}

func flipClick(data *Data, card *tables.Card) {
	cardBack := (*data.StudyData.Flds)[card.ID].Back

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewLabel(cardBack),
			),
			layout.NewSpacer(),
			gradingButtons(data, card),
		),
	)
}

func gradingButtons(data *Data, card *tables.Card) *fyne.Container {
	return container.NewCenter(
		container.NewHBox(
			widget.NewButton("Again", func() { gradeClick(data, card, tables.AGAIN) }),
			widget.NewButton("Hard", func() { gradeClick(data, card, tables.HARD) }),
			widget.NewButton("Good", func() { gradeClick(data, card, tables.GOOD) }),
			widget.NewButton("Easy", func() { gradeClick(data, card, tables.EASY) }),
		),
	)
}
