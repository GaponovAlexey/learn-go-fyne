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
	title := widget.NewLabel("CHOOSE THE OptiOns")
	ch := widget.NewCheckGroup(
		[]string{
			"Options 1",
			"Options 2",
			"Options 3",
			"Options 4",
			"Options 5",
			"Options 6",
			"Options 7",
			"Options 8",
			"Options 9",
			"Options 10",
		},
		nil,
	)

	scrol := container.NewVScroll(
		container.NewVBox(
			ch,
		),
	)
	scrol.SetMinSize(fyne.NewSize(400, 150))
	// res
	res := widget.NewLabel("")
	//but
	btn := widget.NewButton("заказать", func() {
		res.SetText("")
		for _, v := range ch.Selected {
			res.SetText(res.Text + v + "\n")
		}
	})

	w.SetContent(container.NewVBox(title, scrol, btn, res))

	w.ShowAndRun()
	a.Run()
}
