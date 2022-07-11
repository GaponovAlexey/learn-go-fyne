package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(400, 500))

	
	Vbox := container.NewVBox()

	w.SetContent(
		Vbox,
	)

	w.ShowAndRun()
}
