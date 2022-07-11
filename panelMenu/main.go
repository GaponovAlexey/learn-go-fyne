package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello from GO")
	ico, _ := fyne.LoadResourceFromPath("f.ico")
	w.SetIcon(ico)
	w.Resize(fyne.NewSize(500, 400))

	// menu -> second menu -> buttons
	file_item := fyne.NewMenuItem("New File", func() { os.Create("Create.txt") })
	file_item2 := fyne.NewMenuItem("Save", func() { fmt.Println("save") })
	menu1 := fyne.NewMenu("nameMeny", file_item, file_item2)

	//two menu
	actionos_item1 := fyne.NewMenuItem("м1", func() { fmt.Println("м1") })
	actionos_item2 := fyne.NewMenuItem("м2", func() { fmt.Println("м2") })
	actionos_item3 := fyne.NewMenuItem("м3", func() { fmt.Println("м3") })
	menu2 := fyne.NewMenu("Actions", actionos_item1, actionos_item2, actionos_item3)

	//three menu
	info_menu1 := fyne.NewMenuItem("info_menu1", func() { fmt.Println("info_menu1") })
	info_menu1.ChildMenu = fyne.NewMenu(
		"child",
		fyne.NewMenuItem("Print", func() { fmt.Println("Print")	}),
	)
	menu3 := fyne.NewMenu("Format", info_menu1,)


	mainMenu := fyne.NewMainMenu(menu1, menu2, menu3)
	w.SetMainMenu(mainMenu)

	w.ShowAndRun()
}
