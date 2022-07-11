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
	res1 := widget.NewLabel("BUTTOn")
	res2 := widget.NewLabel("BUTTOn")
	res3 := widget.NewLabel("BUTTOn")
	res4 := widget.NewLabel("BUTTOn")

	btn1 := widget.NewButton("green", func() {})
	btn1.Resize(fyne.NewSize(100,40))
	btn2 := widget.NewButton("click2", func() {})
	btn3 := widget.NewButton("click3", func() {})
	btn4 := widget.NewButton("click4", func() {})

	btn_box := container.NewHBox( btn1, btn2, btn3, btn4)
	label_box := container.NewVBox(res1, res2,res3, res4)

	content:= container.NewHBox(btn_box, label_box)
	w.SetContent(
		content,
	)

	w.ShowAndRun()
}
