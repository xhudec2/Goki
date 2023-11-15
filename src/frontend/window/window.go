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

	title := container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel("Decks"),
		layout.NewSpacer(),
	)
	appWindow.SetContent(
		container.NewVBox(
			MenuButtons(),
			layout.NewSpacer(),
			title,
			Decks(),
			layout.NewSpacer(),
		),
	)

	appWindow.CenterOnScreen()
	appWindow.ShowAndRun()
}

func Decks() *fyne.Container {

	canv := canvas.NewRectangle(color.RGBA{85, 85, 85, 255})
	canv.SetMinSize(fyne.NewSize(400, 300))

	stack := container.NewStack(
		canv,
		decksList(),
	)

	return container.NewHBox(
		layout.NewSpacer(),
		stack,
		layout.NewSpacer(),
	)

}

func decksList() *fyne.Container {
	canv := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	canv.SetMinSize(fyne.NewSize(100, 100))

	var data = []string{
		"list", "list", "list", "list", "list", "list",
		"list", "list", "list", "list", "list", "list",
	}
	//	list := widget.NewList(
	//		func() int {
	//			return len(data)
	//		},
	//		func() fyne.CanvasObject {
	//			return widget.NewLabel("template")
	//		},
	//		func(i widget.ListItemID, o fyne.CanvasObject) {
	//			o.(*widget.Label).SetText(data[i])
	//		})
	//
	//	list.Resize(fyne.NewSize(100, 100))
	buttons := container.NewVBox()
	for _, label := range data {
		button := widget.NewButton(label, nil)
		buttons.Add(button)
	}
	list := container.NewScroll(buttons)
	list.SetMinSize(fyne.NewSize(200, 100))
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
	addDeckButton := widget.NewButton("Add Deck", nil)

	return container.NewHBox(
		layout.NewSpacer(),
		decksButton,
		addDeckButton,
		cardsButton,
		layout.NewSpacer(),
	)
}
