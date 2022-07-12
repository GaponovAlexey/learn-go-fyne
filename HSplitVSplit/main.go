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
	w.Resize(fyne.NewSize(500, 400))
	// HSplit
	// NewVSplit
	title:= widget.NewLabel("Орфографи")
	rightLabe:= widget.NewLabel("Правильно")
	incorect:= widget.NewLabel("НеПравильно")
	
	r_text:= widget.NewLabel("Aliqua elit labore non non culpa elit commodo exercitation eiusmod. Deserunt qui est pariatur ullamco minim. Ea officia consequat cupidatat aliquip enim reprehenderit nisi. Aliqua duis officia irure commodo exercitation dolor veniam excepteur culpa deserunt cupidatat aute. Non ea sit deserunt proident aliquip ea eu velit irure pariatur quis ad.")
	r_text.Wrapping = fyne.TextWrapBreak
	
	i_text:= widget.NewLabel("Officia minim officia pariatur reprehenderit quis mollit. Anim tempor adipisicing labore pariatur labore ad magna irure nulla reprehenderit mollit. Nulla aliquip officia ut non eiusmod ullamco. Id ea elit elit labore est id laborum mollit irure incididunt pariatur exercitation elit. Ut amet eiusmod amet ipsum commodo.")
	i_text.Wrapping = fyne.TextWrapBreak
	
	r_box:= container.NewVBox(rightLabe, r_text)
	i_box:= container.NewVBox(incorect, i_text)
	i_box.

	split:= container.NewHSplit(r_box,i_box)
	cont :=  container.NewVBox(title, split)
	
	w.SetContent(cont)
	w.ShowAndRun()
}
