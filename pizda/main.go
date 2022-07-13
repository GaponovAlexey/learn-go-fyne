package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("тiкай с городу")
	w.Resize(fyne.NewSize(100, 100))
	w.SetContent(widget.NewLabel("тоbi пизда "))
	w.Show()
	cont:= container.NewVBox(widget.NewLabel("тоbi пизда "), widget.NewLabel("тiкай с городу"))
	w2 := a.NewWindow("Tobi")
	w2.SetContent(cont)
	w2.Resize(fyne.NewSize(200, 100))
	w2.Show()

	w3 := a.NewWindow("Tobi")
	w3.SetContent(cont)
	w3.Resize(fyne.NewSize(200, 100))
	w3.Show()

	w4 := a.NewWindow("Tobi")
	w4.SetContent(cont)
	w4.Resize(fyne.NewSize(200, 100))
	w4.Show()

	w5 := a.NewWindow("Tobi")
	w5.SetContent(cont)
	w5.Resize(fyne.NewSize(200, 100))
	w5.Show()

	w6 := a.NewWindow("Tobi")
	w6.SetContent(cont)
	w6.Resize(fyne.NewSize(200, 100))
	w6.Show()

	w7 := a.NewWindow("Tobi")
	w7.SetContent(cont)
	w7.Resize(fyne.NewSize(200, 100))
	w7.Show()

	w8 := a.NewWindow("Tobi")
	w8.SetContent(cont)
	w8.Resize(fyne.NewSize(200, 100))
	w8.Show()

	w9 := a.NewWindow("Tobi")
	w9.SetContent(cont)
	w9.Resize(fyne.NewSize(120, 100))
	w9.Show()

	w10 := a.NewWindow("Tobi")
	w10.SetContent(cont)
	w10.Resize(fyne.NewSize(200, 100))
	w10.Show()

	w11 := a.NewWindow("Tobi")
	w11.SetContent(cont)
	w11.Resize(fyne.NewSize(200, 100))
	w11.Show()

	w12 := a.NewWindow("Tobi")
	w12.SetContent(cont)
	w12.Resize(fyne.NewSize(200, 100))
	w12.Show()

	w13 := a.NewWindow("Tobi")
	w13.SetContent(cont)
	w13.Resize(fyne.NewSize(200, 100))
	w13.Show()

	w14 := a.NewWindow("Tobi")
	w14.SetContent(cont)
	w14.Resize(fyne.NewSize(200, 100))
	w14.Show()

	w15 := a.NewWindow("Tobi")
	w15.SetContent(cont)
	w15.Resize(fyne.NewSize(200, 100))
	w15.Show()

	w16 := a.NewWindow("Tobi")
	w16.SetContent(cont)
	w16.Resize(fyne.NewSize(200, 100))
	w16.Show()

	w17 := a.NewWindow("Tobi")
	w17.SetContent(cont)
	w17.Resize(fyne.NewSize(200, 100))
	w17.Show()

	w18 := a.NewWindow("Tobi")
	w18.SetContent(cont)
	w18.Resize(fyne.NewSize(200, 100))
	w18.Show()

	w19 := a.NewWindow("Tobi")
	w19.SetContent(cont)
	w19.Resize(fyne.NewSize(200, 100))
	w19.Show()

	w20 := a.NewWindow("Tobi")
	w20.SetContent(cont)
	w20.Resize(fyne.NewSize(200, 100))
	w20.Show()

	w21 := a.NewWindow("Tobi")
	w21.SetContent(cont)
	w21.Resize(fyne.NewSize(200, 100))
	w21.Show()

	w22 := a.NewWindow("Tobi")
	w22.SetContent(cont)
	w22.Resize(fyne.NewSize(200, 100))
	w22.Show()

	a.Run()
}
