package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(900, 600))
	color:= color.NRGBA{R:123, G:45, B: 120, A: 211}
	label:= canvas.NewText("color", color)


w.SetContent(
	container.NewVBox(
		label,
	))

	w.ShowAndRun()
}
