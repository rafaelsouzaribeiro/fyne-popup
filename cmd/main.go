package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/internal/component/popup"
)

func main() {
	a := app.New()

	w := a.NewWindow("Pop up")
	hello := widget.NewLabel("Demo pop-up")
	showPopUpButton := widget.NewButton("Show Pop-up", nil)
	content := canvas.NewText("This is the content of the pop-up",
		color.Black)

	showPopUpButton.OnTapped = func() {
		popup.ShowPopup("Close", w, content)
	}

	w.Resize(fyne.NewSize(800, 480))
	w.SetContent(container.NewVBox(
		hello,
		showPopUpButton,
	))
	w.ShowAndRun()
}
