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

	channel := make(chan string, 1)

	inputFields, button := inputs(data, w, channel)

	w.SetContent(
		container.NewVBox(
			layout.NewSpacer(),
			deckSelection(data, channel),
			layout.NewSpacer(),
			inputFields,
			layout.NewSpacer(),
			button,
			layout.NewSpacer(),
		),
	)
	w.Show()
}

func inputs(data *Data, window fyne.Window, channel chan string) (*fyne.Container, *widget.Button) {
	frontCont, front := inputWidget("Front")
	backCont, back := inputWidget("Back")

	inputFields := container.NewHBox(
		layout.NewSpacer(),
		frontCont,
		backCont,
		layout.NewSpacer(),
	)
	b := widget.NewButton(
		"Add",
		func() {
			did := tables.GetDeckID(data.StudyData.Decks, <-channel)
			database.AddNewCard(did, front.Text, back.Text, data.StudyData.DB)
			window.Close()
			Draw(data)
		},
	)
	return inputFields, b
}

func deckSelection(data *Data, channel chan string) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Choose Deck"),
		widget.NewSelect(
			tables.DeckNames(data.StudyData.Decks),
			func(s string) {
				channel <- s
			},
		),
	)
}

func inputWidget(label string) (*fyne.Container, *widget.Entry) {
	w := widget.NewEntry()
	w.SetPlaceHolder("Enter text...")
	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(200, 50))
	return container.NewStack(
		bg,
		container.NewGridWithRows(
			2,
			widget.NewLabel(label),
			w,
		),
	), w

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
	cardFrontText := canvas.NewText(cardFront, color.White)
	cardFrontText.TextSize = 20

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				cardFrontText,
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

	cardBackText := canvas.NewText(cardBack, color.White)
	cardBackText.TextSize = 20

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				cardBackText,
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
