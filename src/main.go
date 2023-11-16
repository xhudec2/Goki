package main

import (
	"fmt"
	"src/backend/collation"
	"src/backend/database"
	"src/backend/scheduler"
	"src/backend/tables"
	"src/frontend/window"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	collation.RegisterCollation()
	db, err := database.OpenDB(database.MAIN_DB)
	if err != nil {
		return
	}
	defer database.CloseDB(db)

	cards := make(database.Table[tables.Card], scheduler.Q_SIZE)
	err = database.DBGetter[tables.Card](db, &cards, scheduler.DEFAULT_WHERE)
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	fmt.Println(tables.TodayRelative(db))
	app := app.New()
	app.Settings().SetTheme(&window.Theme{})

	appWindow := app.NewWindow("GOki")
	appWindow.Resize(fyne.NewSize(800, 600))
	appWindow.CenterOnScreen()

	d1 := window.Deck{ID: 1684396573184, Name: "Karaktärsdrag", New: 5, Learn: 10, Due: 15}
	d2 := window.Deck{ID: 1684396927472, Name: "Frisk som en fisk", New: 1, Learn: 1, Due: 1}
	d3 := window.Deck{ID: 1684396937283, Name: "Människokroppen", New: 0, Learn: 0, Due: 0}
	d4 := window.Deck{ID: 1684396946079, Name: "Utseende", New: 0, Learn: 0, Due: 5}
	decks := []window.Deck{
		d1, d2, d3, d4,
	}
	appData := window.AppData{
		App:    app,
		Window: appWindow,
		Decks:  &decks,
		DB:     db,
	}
	window.Draw(&appData)
	appWindow.ShowAndRun()
}
