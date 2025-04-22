// ui package contains the gui version of the pick a gun service
package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
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
	"Blast Crossbow":           resourceBlastCrossbowJPEG,
	"Sawed-off W-870":          resourceSawedOffW870JPEG,
	"Blacktail AC":             resourceBlackTailACJPEG,
	"Punisher MC":              resourcePunisherMCJPEG,
	"Black":                    resourceBlackJPEG,
}

func init() {
	image.RegisterFormat("jpeg", "", jpeg.Decode, jpeg.DecodeConfig)
}

func updateLabels(g *widget.Label, c *widget.Label, p *widget.Label, gi *canvas.Image, sd *core.SaveData) {
	c.SetText(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	g.SetText(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))
	p.SetText(fmt.Sprintf("Selected Character: %s", sd.SelectedCharacter))
	newImage := canvas.NewImageFromResource(gunMap[sd.CurrentGun])
	gi.Image = newImage.Image
	gi.File = newImage.File
	gi.Resource = newImage.Resource
	gi.Refresh()
}

// TODO: Instead of this helper function, see if core handle it instead.
// FileExists checks if the file at path exists and is not a directory
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	a := app.New()
	w := a.NewWindow("re4-pick-a-gun service")

	resolution := fyne.Size{Width: 300, Height: 400}
	w.Resize(resolution)
	w.CenterOnScreen()

	var sd core.SaveData
	characterLabel := widget.NewLabel(fmt.Sprintf("Selected Character: %s", sd.SelectedCharacter))
	chapLabel := widget.NewLabel(fmt.Sprintf("Current Chapter: %d", sd.CurrentChapter))
	gunLabel := widget.NewLabel(fmt.Sprintf("Current Gun: %s", sd.CurrentGun))

	gunImage := canvas.NewImageFromResource(gunMap["Black"])

	gunImage.Resize(fyne.Size{Width: 200, Height: 200})
	gunImage.FillMode = canvas.ImageFillOriginal

	adaButton := widget.NewButton("Seperate Ways(Ada)", func() {
		err := sd.StartGame("A", core.AdaHandguns, core.AdaShotguns, core.AdaRifles, core.AdaSubs, core.AdaSpecials)
		if err != nil {
			dialog.NewError(err, w).Show()
		}
		updateLabels(gunLabel, chapLabel, characterLabel, gunImage, &sd)
	})

	leonButton := widget.NewButton("Main Game(Leon)", func() {
		err := sd.StartGame("L", core.Handguns, core.Shotguns, core.Rifles, core.Subs, core.Magnums)
		if err != nil {
			dialog.NewError(err, w).Show()
		}
		updateLabels(gunLabel, chapLabel, characterLabel, gunImage, &sd)
	})

	confirmSave := dialog.NewConfirm("Saving...", "Confirm save?", func(b bool) {
		if b && sd.CurrentChapter != 0 {
			sd.SaveGame()
		} else {
			dialog.NewError(errors.New("Cannot save without starting the game."), w).Show()
		}
	}, w)

	saveButton := widget.NewButton("save", func() {
		log.Println("saving...")
		if sd.CurrentChapter == 0 {
			dialog.NewError(errors.New("Cannot save without starting the game."), w).Show()
		} else {
			confirmSave.Show()
		}
	})
	loadButton := widget.NewButton("load previously saved data", func() {
		log.Println("loading...")
		err := sd.LoadGame()
		if err != nil {
			dialog.NewError(err, w).Show()
		} else {
			updateLabels(gunLabel, chapLabel, characterLabel, gunImage, &sd)
		}
	})
	rollButton := widget.NewButton("roll", func() {
		log.Println("rolling...")
		if sd.CurrentChapter != sd.FinalChapter {
			sd.RollGun()
			updateLabels(gunLabel, chapLabel, characterLabel, gunImage, &sd)
		} else {
			dialog.NewError(errors.New("Invalid, choose a new playthrough"), w).Show()

		}
	})
	quitButton := widget.NewButton("quit", func() {
		os.Exit(0)
	})
	pickPlayer := container.NewVBox(adaButton, leonButton)
	textBoxes := container.NewVBox(characterLabel, chapLabel, gunLabel, gunImage)
	buttonBox := container.NewVBox(rollButton, loadButton, saveButton, quitButton, textBoxes)

	mainMenu := container.NewVBox(pickPlayer, buttonBox)
	w.SetContent(mainMenu)

	w.ShowAndRun()
}
