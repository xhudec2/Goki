package window

import (
	"fmt"
	"image/color"

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

type AppData struct {
	App    fyne.App
	Window fyne.Window
	Decks  *[]Deck
}

func Draw(data *AppData) {
	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			Decks(data),
			layout.NewSpacer(),
		),
	)
	data.Window.Show()
}

func Decks(data *AppData) *fyne.Container {

	canv := canvas.NewRectangle(color.RGBA{65, 65, 65, 255})
	canv.CornerRadius = 15
	canv.SetMinSize(fyne.NewSize(400, 300))

	stack := container.NewStack(
		canv,
		container.NewPadded(
			container.NewVBox(
				title(),
				decksList(data),
			),
		),
	)

	return container.NewCenter(
		stack,
	)
}

func title() *fyne.Container {
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
		buttons.Add(makeDeckButton(data, index, background))
	}

	list := container.NewScroll(buttons)
	list.SetMinSize(fyne.NewSize(400, 250))
	padded := container.NewPadded(list)
	return container.NewHBox(
		layout.NewSpacer(),
		padded,
		layout.NewSpacer(),
	)
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

func CreateDeck(app fyne.App) {
	w := app.NewWindow("Create Deck")
	w.Resize(fyne.NewSize(600, 450))
	w.CenterOnScreen()
	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Create Deck"),
		),
	)
	w.Show()
}

func DrawDeck(name string) *fyne.Container {
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(name),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}

func StudyDeck(data *AppData, index int) {
	cont := container.NewGridWithRows(
		3,
		container.NewHBox(
			widget.NewLabel("New"),
			layout.NewSpacer(),
			widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].New)),
		),
		container.NewHBox(
			widget.NewLabel("Learn"),
			layout.NewSpacer(),
			widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].Learn)),
		),
		container.NewHBox(
			widget.NewLabel("Due"),
			layout.NewSpacer(),
			widget.NewLabel(fmt.Sprintf("%d", (*data.Decks)[index].Due)),
		),
	)
	canv := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	canv.SetMinSize(fyne.NewSize(300, 10))

	data.Window.SetContent(
		container.NewVBox(
			MenuButtons(data),
			layout.NewSpacer(),
			container.NewCenter(
				container.NewStack(
					canv,
					container.NewHBox(
						cont,
						layout.NewSpacer(),
						DrawDeck((*data.Decks)[index].Name),
					),
				),
			),
			layout.NewSpacer(),
			container.NewCenter(
				widget.NewButton("Study Deck", nil),
			),
			layout.NewSpacer(),
		),
	)
	data.Window.Show()
}
