package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello from GO")
	ico, _ := fyne.LoadResourceFromPath("f.ico")
	w.SetIcon(ico)
	w.Resize(fyne.NewSize(300, 400))

	//VScroll
	//HScroll
	img1 := canvas.NewImageFromFile("g1.jpg")
	img2 := canvas.NewImageFromFile("g2.jpg")
	img3 := canvas.NewImageFromFile("g3.jpg")
	// rec1 := canvas.NewRectangle(color.RGBA{0, 22, 55, 255})
	// rec2 := canvas.NewRectangle(color.RGBA{0, 32, 45, 255})
	// rec3 := canvas.NewRectangle(color.RGBA{0, 42, 25, 255})
	img1.SetMinSize(fyne.NewSize(400, 300))
	img2.SetMinSize(fyne.NewSize(400, 300))
	img3.SetMinSize(fyne.NewSize(400, 300))


	scr := container.NewHScroll(
		container.NewHBox(
			img1,
			img2,
			img3,
		),
	)

	w.SetContent(scr)

	w.ShowAndRun()
	a.Run()
}
