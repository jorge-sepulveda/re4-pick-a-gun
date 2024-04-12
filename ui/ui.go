// ui package contains the gui version of the pick a gun service
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
	"github.com/jorge-sepulveda/re4-pick-a-gun/ui/gun_menu"
	"github.com/jorge-sepulveda/re4-pick-a-gun/ui/start_menu"

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

func main() {
	a := app.New()
	w := a.NewWindow("re4-pick-a-gun service")

	resolution := fyne.Size{Width: 300, Height: 400}
	w.Resize(resolution)
	w.CenterOnScreen()
	var sd core.SaveData

	//Initialize start menuo
	loadButton := widget.NewButton("load previously saved data", nil)
	startButton := widget.NewButton("Start game with options", nil)
	quitButton := widget.NewButton("quit", func() {
		os.Exit(0)
	})

	widgetBox := container.NewVBox(start_menu.GetWidgets(&sd)...)
	widgetBox.Add(loadButton)
	widgetBox.Add(startButton)
	widgetBox.Add(quitButton)

	loadGunsBox := func() {
		widgetBox.Hide()
		gunBox := container.NewVBox(gun_menu.LoadGunMenu(&sd, gunMap, &w)...)
		gunBox.Add(loadButton)
		gunBox.Add(quitButton)
		w.SetContent(gunBox)
	}

	startButton.OnTapped = func() {
		log.Println("Starting Game!")
		err := sd.StartGame()
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}
		loadGunsBox()
	}

	loadButton.OnTapped = func() {
		log.Println("loading...")
		err := sd.LoadGame()
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}
		loadGunsBox()
	}

	w.SetContent(widgetBox)
	w.ShowAndRun()

}
