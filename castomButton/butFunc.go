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
	img := canvas.NewImageFromFile("1.png")

	but1 := createColorButton(color.RGBA{10, 129, 160, 12},
		widget.NewButton("but", func() {}))
	but2 := createImgButton(img, widget.NewButton("but", func() {}))
	
	w.SetContent(container.NewVBox(
		but1,
		but2,
	))
	w.ShowAndRun()
}

func createColorButton(c color.RGBA, b *widget.Button) *fyne.Container {
	button := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(c),
		b,
	)
	return button
}

func createImgButton(img *canvas.Image, b *widget.Button) *fyne.Container {
	button := container.New(
		layout.NewMaxLayout(),
		img,
		b,
	)
	return button
}
