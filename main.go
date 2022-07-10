package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("fyne")

	a := app.New()
	w := a.NewWindow("Hello from GO")
	label := widget.NewLabel("Hello from app")
	label2 := widget.NewLabel("label2")
	w.SetContent(container.NewVBox(label,label2))
	w.ShowAndRun()
}
