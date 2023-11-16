package window

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

type Deck struct {
	ID    int
	Name  string
	New   int
	Learn int
	Due   int
}

type AppData struct {
	App    fyne.App
	Window fyne.Window
	Decks  *[]Deck
	DB     *gorm.DB
}

func Draw(data *AppData) {
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

func MenuButtons(data *AppData) *fyne.Container {
	decksButton := widget.NewButton("Decks", func() { Draw(data) })
	cardsButton := widget.NewButton("Add Card", func() { AddCard(data.App) })
	addDeckButton := widget.NewButton("Create Deck", func() { CreateDeck(data.App) })
	return container.NewHBox(
		layout.NewSpacer(),
		container.NewGridWithColumns(
			3,
			decksButton,
			addDeckButton,
			cardsButton,
		),
		layout.NewSpacer(),
	)
}

func decksMenu(data *AppData) *fyne.Container {

	canv := canvas.NewRectangle(color.RGBA{65, 65, 65, 255})
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

func decksList(data *AppData) *fyne.Container {

	buttons := container.NewVBox()
	for index := range *data.Decks {
		canv := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
		canv.SetMinSize(fyne.NewSize(200, 10))
		buttons.Add(deckListEntry(data, index, canv))
	}

	list := container.NewScroll(buttons)
	list.SetMinSize(fyne.NewSize(400, 250))

	return container.NewHBox(
		layout.NewSpacer(),
		container.NewPadded(list),
		layout.NewSpacer(),
	)
}

func deckListEntry(data *AppData, index int, canv *canvas.Rectangle) *fyne.Container {
	cardData := container.NewWithoutLayout(
		widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].New)),
		widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].Learn)),
		widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].Due)),
	)
	cardData.Objects[0].Move(fyne.NewPos(40, 0))
	cardData.Objects[1].Move(fyne.NewPos(105, 0))
	cardData.Objects[2].Move(fyne.NewPos(160, 0))

	background := container.NewStack(
		canv,
		cardData,
	)
	return makeDeckButton(data, index, background)
}

func makeDeckButton(data *AppData, index int, background *fyne.Container) *fyne.Container {
	return container.NewHBox(
		widget.NewButton(
			(*data.Decks)[index].Name,
			func() {
				StudyDeck(data, index)
			},
		),
		layout.NewSpacer(),
		background,
	)
}
