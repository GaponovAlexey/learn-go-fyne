package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(900, 600))

	// img2
	circl2 := canvas.NewCircle(color.Black)
	//рамка
	circl2.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	circl2.StrokeWidth = 5 // рамка цвет
	// квадрат
	rectangle := canvas.NewRectangle(color.NRGBA{R: 77, G: 55, B: 33, A: 255})
	// линия
	line:= canvas.NewLine(color.NRGBA64{R: 122, G: 33, B: 56, A: 255})
	line.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	line.StrokeWidth = 5 // рамка цвет
	//icon
	icon:= widget.NewIcon(theme.FileIcon())

	cont1 := container.NewGridWithColumns(2,rectangle,icon)
	cont2 := container.NewGridWithColumns(2,line, circl2)

	content := container.NewGridWithRows(2, cont1, cont2)
	w.SetContent(
		content,
	)

	w.ShowAndRun()
}
