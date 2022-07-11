package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(400, 500))
	res := widget.NewLabel("BUTTOn")
	btn1 := widget.NewButton("green", func() {
		})
	btn2 := widget.NewButton("click2", func() {})
	btn3 := widget.NewButton("click3", func() {})
	btn4 := widget.NewButton("click4", func() {})

	Vbox := container.NewVBox(res, btn1, btn2, btn3, btn4)
	w.SetContent(
		Vbox,
	)

	w.ShowAndRun()
}
