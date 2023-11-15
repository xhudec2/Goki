package window

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Draw() {
	app := app.New()
	app.Settings().SetTheme(&Theme{})

	appWindow := app.NewWindow("GOki")
	appWindow.Resize(fyne.NewSize(800, 600))

	appWindow.SetContent(
		container.NewVBox(
			MenuButtons(),
			layout.NewSpacer(),
			Decks(),
			layout.NewSpacer(),
		),
	)

	appWindow.CenterOnScreen()
	appWindow.ShowAndRun()
}

func Decks() *fyne.Container {

	canv := canvas.NewRectangle(color.RGBA{65, 65, 65, 255})
	canv.CornerRadius = 10
	canv.SetMinSize(fyne.NewSize(400, 300))

	title := container.NewHBox(
		widget.NewLabel("Deck Name"),
		layout.NewSpacer(),
		container.NewGridWithColumns(
			3,
			widget.NewLabel("New"),
			widget.NewLabel("Learn"),
			widget.NewLabel("Due"),
		),
	)

	stack := container.NewStack(
		canv,
		container.NewPadded(
			container.NewVBox(
				title,
				decksList(),
			),
		),
	)

	return container.NewCenter(
		stack,
	)
}

func decksList() *fyne.Container {

	var data = []string{
		"list", "list", "list", "list", "list", // "list",
		//"list", "list", "list", "list", "list", "list",
	}
	buttons := container.NewVBox()
	for _, label := range data {
		canv := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
		canv.SetMinSize(fyne.NewSize(200, 10))

		cardData := container.NewWithoutLayout(
			widget.NewLabel("0"),
			widget.NewLabel("0"),
			widget.NewLabel("0"),
		)
		cardData.Objects[0].Move(fyne.NewPos(40, 0))
		cardData.Objects[1].Move(fyne.NewPos(110, 0))
		cardData.Objects[2].Move(fyne.NewPos(165, 0))
		bg := container.NewStack(
			canv,
			cardData,
		)
		deck := container.NewHBox(
			widget.NewButton(label, nil),
			layout.NewSpacer(),
			bg,
		)
		buttons.Add(deck)
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

func MenuButtons() *fyne.Container {

	decksButton := widget.NewButton("Decks", nil)
	cardsButton := widget.NewButton("Add Card", nil)
	addDeckButton := widget.NewButton("Create Deck", nil)

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
