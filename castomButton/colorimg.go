package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	but := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.RGBA{R: 0, G: 133, B: 133, A: 255}),
		widget.NewButton("but", func() {}),
	)
	but2 := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.RGBA{70,129,60,1}),
		widget.NewButton("but", func() {}),
	)
	img:= canvas.NewImageFromFile("1.png")
	but3 := container.New(
		layout.NewMaxLayout(),
		img,
		widget.NewButton("but", func() {}),
	)

	w.SetContent(container.NewVBox(but, but2, but3))
	w.ShowAndRun()
}
