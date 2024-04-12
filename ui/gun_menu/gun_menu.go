package gun_menu

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func LoadGunMenu(sd *core.SaveData, gunPool map[string]*fyne.StaticResource, w *fyne.Window) []fyne.CanvasObject {
	var gunWidgets = []fyne.CanvasObject{}
	chapLabel := widget.NewLabel(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	gunLabel := widget.NewLabel(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))
	gunImage := canvas.NewImageFromResource(gunPool[sd.CurrentGun])
	updateLabels(gunLabel, chapLabel, gunImage, sd, gunPool)

	rollButton := widget.NewButton("roll", func() {
		//log.Println("rolling...")
		err := sd.RollGun()
		if err != nil {
			dialog.NewError(err, *w).Show()
			return
		}
		updateLabels(gunLabel, chapLabel, gunImage, sd, gunPool)

	})

	gunImage.Resize(fyne.Size{Width: 200, Height: 200})
	gunImage.FillMode = canvas.ImageFillOriginal

	confirmSave := dialog.NewConfirm("Saving...", "Confirm save?", func(b bool) {
		if b {
			sd.SaveGame()
		}
	}, *w)
	saveButton := widget.NewButton("save", func() {
		fmt.Println("saving...")
		confirmSave.Show()
	})

	gunWidgets = append(gunWidgets, chapLabel, gunLabel, gunImage, rollButton, saveButton)
	return gunWidgets
}

func updateLabels(g *widget.Label, c *widget.Label, gi *canvas.Image, sd *core.SaveData, gunPool map[string]*fyne.StaticResource) {
	c.SetText(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	g.SetText(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))
	newImage := canvas.NewImageFromResource(gunPool[sd.CurrentGun])
	gi.Image = newImage.Image
	gi.File = newImage.File
	gi.Resource = newImage.Resource
	gi.Refresh()
}
