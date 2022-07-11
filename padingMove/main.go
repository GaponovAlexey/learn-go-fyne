package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(400, 400))

	btn := widget.NewButton("Click", func() {	})
	btn.Resize(fyne.NewSize(200,20))
	btn.Move(fyne.NewPos(10,50))

	btn2 := widget.NewButton("Click", func() {	})
	btn2.Resize(fyne.NewSize(100,20))
	btn2.Move(fyne.NewPos(10,30))


	cont:= container.NewWithoutLayout(btn, btn2)
	w.SetContent(
		cont,
	)


	w.ShowAndRun()
}
