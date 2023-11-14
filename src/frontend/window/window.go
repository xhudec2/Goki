package window

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func DrawWindow() {
	GOki := app.New()
	appWindow := GOki.NewWindow("GOki")

	appWindow.Resize(fyne.NewSize(800, 600))

	foo := func() {
		appWindow.SetTitle("Clicked")
		time.Sleep(1 * time.Second)
		appWindow.SetTitle("GOki")
	}
	decksButton := widget.NewButton("Decks", foo)
	cardsButton := widget.NewButton("Add Card", foo)
	content := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), decksButton, cardsButton, layout.NewSpacer())
	appWindow.SetContent(container.New(layout.NewVBoxLayout(), content))
	appWindow.ShowAndRun()
}
