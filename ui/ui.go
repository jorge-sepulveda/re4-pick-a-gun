// ui package contains the gui version of the pick a gun service
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func init() {
	image.RegisterFormat("jpeg", "", jpeg.Decode, jpeg.DecodeConfig)
}

func updateLabels(g *widget.Label, c *widget.Label, gi *canvas.Image, sd *core.SaveData) {
	c.SetText(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	g.SetText(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))
	newImage := canvas.NewImageFromFile("./img/" + strings.ReplaceAll(sd.CurrentGun, " ", "_") + ".JPEG")
	gi.Image = newImage.Image
	gi.File = newImage.File
	gi.Resource = newImage.Resource
	gi.Refresh()
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello World")

	resolution := fyne.Size{Width: 300, Height: 400}
	w.Resize(resolution)
	w.CenterOnScreen()

	var sd core.SaveData
	err := sd.StartGame(core.Handguns, core.Shotguns, core.Rifles, core.Subs, core.Magnums)
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
	chapLabel := widget.NewLabel(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	gunLabel := widget.NewLabel(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))

	gunImage := canvas.NewImageFromFile("./img/" + strings.ReplaceAll(sd.CurrentGun, " ", "_") + ".JPEG")

	gunImage.Resize(fyne.Size{Width: 200, Height: 200})
	gunImage.FillMode = canvas.ImageFillOriginal
	if err != nil {
		fmt.Println(err)
	}
	confirmSave := dialog.NewConfirm("Saving...", "Confirm save?", func(b bool) {
		if b {
			sd.SaveGame()
		}
	}, w)

	saveButton := widget.NewButton("save", func() {
		log.Println("saving...")
		confirmSave.Show()
	})
	loadButton := widget.NewButton("load previously saved data", func() {
		log.Println("loading...")
		sd.LoadGame()
		updateLabels(gunLabel, chapLabel, gunImage, &sd)
	})
	rollButton := widget.NewButton("roll", func() {
		log.Println("rolling...")
		if sd.CurrentChapter != core.MAXCHAPTER {
			sd.RollGun()
			updateLabels(gunLabel, chapLabel, gunImage, &sd)
		} else {
			fmt.Println("No more chapters, Stranger")
		}
	})
	quitButton := widget.NewButton("quit", func() {
		os.Exit(0)
	})

	textBoxes := container.NewVBox(chapLabel, gunLabel, gunImage)
	buttonBox := container.NewVBox(rollButton, loadButton, saveButton, quitButton, textBoxes)
	w.SetContent(buttonBox)
	w.ShowAndRun()
}
