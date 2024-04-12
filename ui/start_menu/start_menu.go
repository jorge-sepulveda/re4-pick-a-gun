package start_menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func GetWidgets(sd *core.SaveData) []fyne.CanvasObject {
	checks := []fyne.CanvasObject{}

	hg := widget.NewCheck("Handguns", func(b bool) {
		sd.SetHandguns(b)
	})

	sg := widget.NewCheck("Shotguns", func(b bool) {
		sd.SetShotguns(b)
	})

	rf := widget.NewCheck("Rifles", func(b bool) {
		sd.SetRifles(b)
	})

	mn := widget.NewCheck("Magnums", func(b bool) {
		sd.SetMagnums(b)
	})

	sb := widget.NewCheck("Submachine Guns", func(b bool) {
		sd.SetSubs(b)
	})

	sp := widget.NewCheck("Specials", func(b bool) {
		sd.SetSpecials(b)
	})

	checks = append(checks, hg, sg, rf, mn, sb, sp)
	return checks
}
