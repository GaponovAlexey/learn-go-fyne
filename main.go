package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello from GO")
	ico, _ := fyne.LoadResourceFromPath("f.ico")
	w.SetIcon(ico)
	w.Resize(fyne.NewSize(500, 400))

	//accordion
	item1 := widget.NewAccordionItem("name", widget.NewLabel("inf"))
	item2 := widget.NewAccordionItem("name", widget.NewLabel(""))
	acordion := widget.NewAccordion(item1,item2)

	w.SetContent(container.NewHBox(acordion))
	w.ShowAndRun()
}
