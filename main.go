package main

import (
	"fmt"
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
	ic, _ := fyne.LoadResourceFromPath("f.ico")
	w.SetIcon(ic)
	w.Resize(fyne.NewSize(500, 400))

	label := widget.NewLabel("Регистрация")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(200, 2))
	//entry
	name := widget.NewEntry()
	name.Resize(fyne.NewSize(300, 40))
	name.Move(fyne.NewPos(100, 60))
	name.SetPlaceHolder("Name")

	pass := widget.NewPasswordEntry()
	pass.Resize(fyne.NewSize(300, 40))
	pass.Move(fyne.NewPos(100, 120))
	pass.SetPlaceHolder("Пароль")

	email := widget.NewEntry()
	email.Resize(fyne.NewSize(300, 40))
	email.Move(fyne.NewPos(100, 180))
	email.SetPlaceHolder("email")

	dopInfo := widget.NewEntry()
	dopInfo.Resize(fyne.NewSize(300, 40))
	dopInfo.Wrapping = fyne.TextWrapBreak
	dopInfo.Move(fyne.NewPos(100, 240))
	dopInfo.SetPlaceHolder("Дополнительная информация")
	//btn
	btn := widget.NewButton("Зарегистрироваться", func() {
		file, _ := os.Create("file/user.txt")
		defer file.Close()
		data := fmt.Sprintf("name:%v\npass:%v\nemail:%v\ndopInfo:%v\n", name.Text, pass.Text, email.Text, dopInfo.Text)
		file.WriteString(data)
	})
	btn.Resize(fyne.NewSize(200, 40))
	btn.Move(fyne.NewPos(150, 300))
	//cont
	cont := container.NewWithoutLayout(label, name, pass, email, dopInfo, btn)

	w.SetContent(
		cont,
	)

	w.ShowAndRun()
}
