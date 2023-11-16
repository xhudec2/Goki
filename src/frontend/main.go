package main

import (
	. "src/frontend/window"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(&Theme{})

	appWindow := app.NewWindow("GOki")
	appWindow.Resize(fyne.NewSize(800, 600))
	appWindow.CenterOnScreen()

	//d1 := Deck{ID: 1, Name: "Karaktärsdrag", New: 5, Learn: 10, Due: 15}
	//d2 := Deck{ID: 2, Name: "Frisk som en fisk", New: 1, Learn: 1, Due: 1}
	//d3 := Deck{ID: 3, Name: "Människokroppen", New: 0, Learn: 0, Due: 0}
	//d4 := Deck{ID: 4, Name: "Utseende", New: 0, Learn: 0, Due: 5}
	d1 := Deck{ID: 1, Name: "Deck 1", New: 5, Learn: 10, Due: 15}
	d2 := Deck{ID: 2, Name: "Deck 2", New: 1, Learn: 1, Due: 1}
	d3 := Deck{ID: 3, Name: "Deck 3", New: 0, Learn: 0, Due: 0}
	d4 := Deck{ID: 4, Name: "Deck 4", New: 0, Learn: 0, Due: 5}
	decks := []Deck{
		d1, d2, d3, d4,
	}
	appData := AppData{
		App:    app,
		Window: appWindow,
		Decks:  &decks,
	}
	Draw(&appData)
	appWindow.ShowAndRun()
}
