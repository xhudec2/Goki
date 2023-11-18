package window

import (
	"fmt"
	"image/color"
	"src/backend/database"
	"src/backend/scheduler"
	"src/backend/tables"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Deck struct {
	ID    int
	Name  string
	New   int
	Learn int
	Due   int
}

type Data struct {
	App       fyne.App
	Window    fyne.Window
	StudyData *scheduler.StudyData
}

func Draw(data *Data) {
	decks := make(tables.Decks, 10)
	err := database.ParseDecks(data.StudyData.DB, &decks)
	if err != nil {
		return
	}
	data.StudyData.Decks = &decks
	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			decksMenu(data),
			layout.NewSpacer(),
		),
	)
	data.Window.Show()
}

func MenuButtons(data *Data) *fyne.Container {
	decksButton := widget.NewButton("Decks", func() { Draw(data) })
	cardsButton := widget.NewButton("Add Card", func() { AddCard(data) })
	addDeckButton := widget.NewButton("Create Deck", func() { CreateDeck(data) })
	delDeckButton := widget.NewButton("Delete Deck", func() { DeleteDeck(data) })
	return container.NewHBox(
		layout.NewSpacer(),
		container.NewGridWithColumns(
			4,
			decksButton,
			addDeckButton,
			cardsButton,
			delDeckButton,
		),
		layout.NewSpacer(),
	)
}

func decksMenu(data *Data) *fyne.Container {

	canv := canvas.NewRectangle(color.RGBA{70, 70, 70, 255})
	canv.CornerRadius = 15
	canv.SetMinSize(fyne.NewSize(400, 300))

	stack := container.NewStack(
		canv,
		container.NewPadded(
			container.NewVBox(
				header(),
				decksList(data),
			),
		),
	)

	return container.NewCenter(
		stack,
	)
}

func header() *fyne.Container {
	return container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel("Deck"),
		layout.NewSpacer(),
		container.NewGridWithColumns(
			3,
			widget.NewLabel("New"),
			widget.NewLabel("Learn"),
			widget.NewLabel("Due"),
		),
	)
}

func decksList(data *Data) *fyne.Container {

	decks := data.StudyData.Decks

	buttons := container.NewVBox()
	for index := range *decks {
		deck := (*decks)[index]
		canv := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
		canv.SetMinSize(fyne.NewSize(200, 10))
		buttons.Add(deckListEntry(data, &deck, canv))
	}

	list := container.NewScroll(buttons)
	list.SetMinSize(fyne.NewSize(400, 250))

	return container.NewHBox(
		layout.NewSpacer(),
		container.NewPadded(list),
		layout.NewSpacer(),
	)
}

func deckListEntry(data *Data, deck *tables.Deck, canv *canvas.Rectangle) *fyne.Container {
	err := database.GetDeckCardData(deck, data.StudyData.DB)
	if err != nil {
		fmt.Println(err)
	}
	cardData := container.NewWithoutLayout(
		widget.NewLabel(fmt.Sprintf("%d", deck.New)),
		widget.NewLabel(fmt.Sprintf("%d", deck.Learn)),
		widget.NewLabel(fmt.Sprintf("%d", deck.Due)),
	)
	cardData.Objects[0].Move(fyne.NewPos(40, 0))
	cardData.Objects[1].Move(fyne.NewPos(105, 0))
	cardData.Objects[2].Move(fyne.NewPos(160, 0))

	background := container.NewStack(
		canv,
		cardData,
	)
	return makeDeckButton(data, deck, background)
}

func makeDeckButton(data *Data, deck *tables.Deck, background *fyne.Container) *fyne.Container {
	return container.NewHBox(
		widget.NewButton(
			deck.Name,
			func() {
				ShowDeckMenu(data, deck)
			},
		),
		layout.NewSpacer(),
		background,
	)
}
