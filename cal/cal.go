package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello from GO")
	label := widget.NewLabel("Ведите первое число")
	imput := widget.NewEntry()

	label2 := widget.NewLabel("Ведите первое число")
	imput2 := widget.NewEntry()

	answer := widget.NewLabel("")
	
	btn := widget.NewButton("equals", func() {
		n1, err := strconv.ParseFloat(imput.Text, 64)
		n2, err := strconv.ParseFloat(imput2.Text, 64)
		if err != nil {
			answer.SetText("err from input")
		} else {
			sum := n1 + n2
			sum2 := n1 - n2
			answer.SetText(fmt.Sprintf("(+)%f\n(-)%f\n", sum, sum2))
		}
	})

	w.SetContent(container.NewVBox(
		label,
		imput,
		label2,
		imput2,
		btn,
		answer,
	))

	w.ShowAndRun()
}
