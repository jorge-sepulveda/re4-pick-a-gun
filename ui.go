// ui package contains the gui version of the pick a gun service
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
	"image"
	"image/jpeg"
	"log"
	"os"
)

var gunMap = map[string]*fyne.StaticResource{
	"SR-09 R":                  resourceInfiniteRocketLauncherJPEG,
	"Punisher":                 resourcePunisherJPEG,
	"Red9":                     resourceRed9JPEG,
	"Blacktail":                resourceBlacktailJPEG,
	"Matilda":                  resourceMatildaJPEG,
	"Sentinel Nine":            resourceSentinelNineJPEG,
	"W-870":                    resourceW870JPEG,
	"Riot Gun":                 resourceRiotGunJPEG,
	"Striker":                  resourceStrikerJPEG,
	"Skull Shaker":             resourceSkullShakerJPEG,
	"SR M1903":                 resourceSRM1903JPEG,
	"Stingray":                 resourceStingrayJPEG,
	"CQBR Assault Rifle":       resourceCQBRAssaultRifleJPEG,
	"Broken Butterfly":         resourceBrokenButterflyJPEG,
	"Killer7":                  resourceKiller7JPEG,
	"TMP":                      resourceTMPJPEG,
	"LE 5":                     resourceLE5JPEG,
	"Handcannon":               resourceHandcannonJPEG,
	"Infinite Rocket Launcher": resourceInfiniteRocketLauncherJPEG,
	"Chicago Sweeper":          resourceChicagoSweeperJPEG,
}

func init() {
	image.RegisterFormat("jpeg", "", jpeg.Decode, jpeg.DecodeConfig)
}

func updateLabels(g *widget.Label, c *widget.Label, gi *canvas.Image, sd *core.SaveData) {
	c.SetText(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	g.SetText(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))
	//	newImage := canvas.NewImageFromFile("./img/" + strings.ReplaceAll(sd.CurrentGun, " ", "_") + ".JPEG")
	newImage := canvas.NewImageFromResource(gunMap[sd.CurrentGun])
	gi.Image = newImage.Image
	gi.File = newImage.File
	gi.Resource = newImage.Resource
	gi.Refresh()
}

func main() {
	a := app.New()

	w := a.NewWindow("re4-pick-a-gun service")

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

	gunImage := canvas.NewImageFromResource(gunMap[sd.CurrentGun])

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
		err = sd.LoadGame()
		if err != nil {
			dialog.NewError(err, w).Show()
		}
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
