package main

import (
	"os"

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
	w.Resize(fyne.NewSize(300, 400))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Text...")
	cont := widget.NewLabel("")
	btn := widget.NewButton("save", func() {
		file,_ := os.Create("info.txt") //создание
		defer file.Close()// закрытие
		file.WriteString(entry.Text) // запись перезапись
		cont.SetText(entry.Text) // запись в переменую
	})
	w.SetContent(
		container.NewVBox(entry, btn, cont),
	)

	w.ShowAndRun()
}
