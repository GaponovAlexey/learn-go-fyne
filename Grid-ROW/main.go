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

	oval1 := canvas.NewCircle(color.NRGBA{R: 155, G: 55, B: 33, A: 255})
	oval2 := canvas.NewCircle(color.NRGBA{R: 155, G: 55, B: 33, A: 255})
	oval3 := canvas.NewCircle(color.NRGBA{R: 155, G: 55, B: 33, A: 255})
	oval4 := canvas.NewCircle(color.NRGBA{R: 155, G: 55, B: 33, A: 255})

	row1 := container.NewGridWithRows(2, oval1, oval2)
	row2 := container.NewGridWithRows(2, oval3, oval4)

	content := container.NewGridWithColumns(2, row1, row2)
	w.SetContent(
		content,
	)

	w.ShowAndRun()
}
