package window

import (
	"fmt"
	"image/color"
	"src/backend/database"
	"src/backend/tables"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateDeck(data *Data) {
	w := data.App.NewWindow("Create Deck")
	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(400, 200))
	bg.CornerRadius = 15

	entryRect := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	entryRect.SetMinSize(fyne.NewSize(300, 0))
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Text...")

	entryStack := container.NewStack(
		entryRect,
		entry,
	)

	s := container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Deck Name"),
			container.NewHBox(
				layout.NewSpacer(),
				entryStack,
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton(
					"Create",
					func() {
						database.AddNewDeck(entry.Text, data.StudyData.DB)
						w.Close()
						Draw(data)
					}),
				widget.NewButton("Close", func() { w.Close() }),
				layout.NewSpacer(),
			),
		),
	)

	w.SetContent(
		container.NewStack(
			bg,
			container.NewCenter(s),
		),
	)
	w.Show()
}

func DeleteDeck(data *Data) {
	w := data.App.NewWindow("Delete Deck")
	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(400, 200))
	bg.CornerRadius = 15

	channel := make(chan string, 1)

	entryRect := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	entryRect.SetMinSize(fyne.NewSize(300, 0))

	entry := widget.NewSelect(tables.DeckNames(data.StudyData.Decks), func(s string) { channel <- s })

	entryStack := container.NewStack(
		entryRect,
		entry,
	)

	s := container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Deck Name"),
			container.NewHBox(
				layout.NewSpacer(),
				entryStack,
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton(
					"Delete",
					func() {
						database.DeleteDeck(<-channel, data.StudyData.DB)
						w.Close()
						Draw(data)
					}),
				widget.NewButton("Close", func() { w.Close() }),
				layout.NewSpacer(),
			),
		),
	)

	w.SetContent(
		container.NewStack(
			bg,
			container.NewCenter(s),
		),
	)
	w.Show()
}

func ShowDeckMenu(data *Data, deck *tables.Deck) {
	var content *fyne.Container
	err := database.GetDeckCardData(deck, data.StudyData.DB)
	if err != nil {
		fmt.Println(err)
	}
	if deck.New+deck.Learn+deck.Due == 0 {
		content = NothingToStudy(data)
	} else {
		content = showDeck(data, deck)
	}
	data.Window.SetContent(content)
	data.Window.Show()
}

func drawDeck(name string) *fyne.Container {
	return container.NewCenter(
		container.NewCenter(widget.NewLabel(name)),
	)
}

func deckMenu(deck *tables.Deck) *fyne.Container {
	labels := []string{
		"New",
		"Learn",
		"Due",
	}
	values := []string{
		fmt.Sprintf("%d", deck.New),
		fmt.Sprintf("%d", deck.Learn),
		fmt.Sprintf("%d", deck.Due),
	}
	deckData := container.NewGridWithRows(3)
	for i := range labels {
		deckData.Add(
			container.NewHBox(
				widget.NewLabel(labels[i]),
				layout.NewSpacer(),
				widget.NewLabel(values[i]),
			),
		)
	}
	canv := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	canv.SetMinSize(fyne.NewSize(300, 10))

	return container.NewStack(
		canv,
		container.NewHBox(
			deckData,
			layout.NewSpacer(),
			drawDeck(deck.Name),
		),
	)
}

func NothingToStudy(data *Data) *fyne.Container {
	return container.NewVBox(
		MenuButtons(data),
		layout.NewSpacer(),
		container.NewCenter(
			widget.NewLabel("Nothing to study! Come back later."),
		),
		container.NewCenter(
			widget.NewButton("Return to main menu", func() { Draw(data) }),
		),
		layout.NewSpacer(),
	)
}

func showDeck(data *Data, deck *tables.Deck) *fyne.Container {
	return container.NewVBox(
		MenuButtons(data),
		layout.NewSpacer(),
		container.NewCenter(deckMenu(deck)),
		layout.NewSpacer(),
		container.NewCenter(widget.NewButton("Study Deck", func() { CycleDeck(data, deck) })),
		layout.NewSpacer(),
	)
}
