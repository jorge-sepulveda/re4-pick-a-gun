package start_menu

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func GetWidgets(sd core.SaveData) []fyne.CanvasObject {
	checks := []fyne.CanvasObject{}
	hg := widget.NewCheck("Handguns", func(b bool) {
		sd.SetHandguns(b)
	})
	checks = append(checks, hg)
	log.Println("WTF")
	return checks
}
