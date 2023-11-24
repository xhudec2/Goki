package main

import (
	"fmt"
	"os"
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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	db, err := database.OpenDB(home + database.MAIN_DB)
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

	appData := window.Data{
		App:    app,
		Window: appWindow,
		StudyData: &scheduler.StudyData{
			DB: db,
		},
	}
	window.Draw(&appData)
	appWindow.ShowAndRun()
}
