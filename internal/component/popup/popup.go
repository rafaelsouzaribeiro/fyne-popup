package popup

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowPopup(w fyne.Window, content *canvas.Text) {
	var popup *widget.PopUp

	popUpContent := container.NewVBox(
		content,
		widget.NewButton("Close", func() {
			popup.Hide()
		}),
	)

	popup = widget.NewModalPopUp(popUpContent, w.Canvas())
	popup.Show()
}
