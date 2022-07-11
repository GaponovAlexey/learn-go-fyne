package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	text := "Deserunt minim aute eiusmod deserunt culpa ipsum pariatur commodo excepteur ipsum excepteur. Consectetur incididunt eiusmod est excepteur eiusmod dolore labore eiusmod nulla irure cupidatat labexcepteur ipsum excepteur. Consectetur incididunt eiusmod est eexcepteur ipsum excepteur. Consectetur incididunt eiusmod est eexcepteur ipsum excepteur. Consectetur incididunt eiusmod est eexcepteur ipsum excepteur. Consectetur incididunt eiusmod est eore. Consequat qui nostrud aute nisi excepteur non sunt. Culpa laborum irure sint sunt excepteur incididunt ad non quis Lorem."
	//accordion
	textLabel := widget.NewLabel(text)
	textLabel.Wrapping = fyne.TextWrapBreak

	item1 := widget.NewAccordionItem("name", textLabel)
	
	item2 := widget.NewAccordionItem("name2",
	 widget.NewAccordion(
		widget.NewAccordionItem(
			"ddd", 
			widget.NewLabel("ypi"),
			),
		),
	)
	acordion := widget.NewAccordion(item1, item2)

	w.SetContent(acordion)
	w.ShowAndRun()
}
