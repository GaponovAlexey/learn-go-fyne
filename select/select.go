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
	ico, _ := fyne.LoadResourceFromPath("f.ico")
	w.SetIcon(ico)
	w.Resize(fyne.NewSize(300, 400))

	//VScroll
	//HScroll
	label := widget.NewLabel("")
	sel := widget.NewSelect(
		[]string{
			"op1",
			"op2",
			"op3",
			"op4",
			"op5",
		},
		nil,
	)
	sel.PlaceHolder = "выбери из списка"
	//btn
	btn:= widget.NewButton("select", func() {
		label.SetText("вы выбрали: " + sel.Selected)
	})

	//config
	w.SetContent(container.NewVBox(sel,btn, label))
	w.ShowAndRun()
	a.Run()
}
