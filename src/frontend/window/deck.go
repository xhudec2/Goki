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

func StudyDeck(data *AppData, index int) {
	deck := (*data.Decks)[index]
	var content *fyne.Container
	if deck.New+deck.Learn+deck.Due == 0 {
		content = nothingToStudy(data)
	} else {
		content = showDeck(data, index)
	}
	data.Window.SetContent(content)
	data.Window.Show()
}

func drawDeck(name string) *fyne.Container {
	return container.NewCenter(
		container.NewCenter(widget.NewLabel(name)),
	)
}

func deckMenu(deck *Deck) *fyne.Container {
	labels := map[string]string{
		"New":   fmt.Sprintf("%d", deck.New),
		"Learn": fmt.Sprintf("%d", deck.Learn),
		"Due":   fmt.Sprintf("%d", deck.Due),
	}
	deckData := container.NewGridWithRows(3)
	for key, value := range labels {
		deckData.Add(
			container.NewHBox(
				widget.NewLabel(key),
				layout.NewSpacer(),
				widget.NewLabel(value),
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

func nothingToStudy(data *AppData) *fyne.Container {
	return container.NewVBox(
		MenuButtons(data),
		layout.NewSpacer(),
		container.NewCenter(
			widget.NewLabel("Nothing to study! Come back later."),
		),
		layout.NewSpacer(),
	)
}

func showDeck(data *AppData, index int) *fyne.Container {
	return container.NewVBox(
		MenuButtons(data),
		layout.NewSpacer(),
		container.NewCenter(deckMenu(&(*data.Decks)[index])),
		layout.NewSpacer(),
		container.NewCenter(widget.NewButton("Study Deck", func() { CycleCards(data, index) })),
		layout.NewSpacer(),
	)
}
