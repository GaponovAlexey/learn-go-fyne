package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	w.Resize(fyne.NewSize(250, 400))
	title := widget.NewLabel("Checkout")

	nameLabe := widget.NewLabel("it's your name:")
	name := widget.NewEntry()

	stolLabel := widget.NewLabel("table number")
	stol := widget.NewEntry()

	foodLabel := widget.NewLabel("Change food")
	food := widget.NewCheckGroup([]string{"beekon", "milk", "fish", "pamela"}, func(s []string) {})
	result := widget.NewLabel("order:")
	btn := widget.NewButton("order", func() {
		//name
		userName := name.Text
		//foods
		var foodsS string
		for _, v := range food.Selected {
			foodsS += " " + v
		}
		//numberStol
		stol := stol.Text

		fmt.Printf("Name client: %s\nMenu: %s\nTable number: %s", userName, foodsS, stol)
		result.SetText(fmt.Sprintf("Name client:%s\nMenu:%s\nTable number:%s", userName, foodsS, stol))
	})
	w.SetContent(container.NewVBox(
		title,
		nameLabe,
		name,
		stolLabel,
		stol,
		foodLabel,
		food,
		btn,
		result,
	))

	w.ShowAndRun()
}
