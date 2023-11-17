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
	err = database.DBGetter[tables.Card](db, &cards, "", "")
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	app := app.New()
	app.Settings().SetTheme(&window.Theme{})

	appWindow := app.NewWindow("GOki")
	appWindow.Resize(fyne.NewSize(800, 600))
	appWindow.CenterOnScreen()

	decks := make(tables.Decks, 10)
	err = database.ParseDecks(db, &decks)
	if err != nil {
		return
	}
	appData := window.Data{
		App:    app,
		Window: appWindow,
		StudyData: &scheduler.StudyData{
			DB:    db,
			Decks: &decks,
		},
	}
	window.Draw(&appData)
	appWindow.ShowAndRun()
}
