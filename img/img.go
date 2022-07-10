package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(900, 600))

	img1 := canvas.NewImageFromFile("img/1.jpg")
	// img2 := canvas.NewImageFromFile("img/2.jpg")

	label := widget.NewLabel("Картинка")

w.SetContent(
	container.NewGridWithColumns(2,img1,label),
)

	w.ShowAndRun()
}
