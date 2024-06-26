// Package cli runs a cli version of the pick-a-gun service. Initially used for debugging but this can also be used as a text only version
package main

import (
	"fmt"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
	"os"
)

func main() {
	var sd core.SaveData
	var option string
	sd.SetHandguns(true)
	sd.SetSubs(true)
	sd.SetShotguns(true)
	sd.SetRifles(true)
	sd.SetMagnums(true)
	err := sd.StartGame()
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
	fmt.Println("Welcome to the RE4 Pick a gun service!")
	fmt.Printf("Starting in Chapter %d\n", sd.CurrentChapter)
	fmt.Printf("Current Gun: %s\n", sd.CurrentGun)
	fmt.Printf("Used guns list %+v\n", sd.UsedGuns)
	fmt.Printf("existing guns list %v\n", sd.GunsList)
	for {
		fmt.Println("Choose your weapon, Stranger: ")
		fmt.Scan(&option)
		switch option {
		case string('q'):
			fmt.Println("Quit command sent.")

			os.Exit(0)
		case string('i'):
			fmt.Println("Current Game info:")
			sd.PrintData()
		case string('l'):
			fmt.Println("Load command sent.")
			err = sd.LoadGame()
			if err != nil {
				print(err.Error() + "\n")
			}
			sd.PrintData()
		case string('s'):
			fmt.Println("Confirm save? [y]")
			fmt.Scan(&option)
			if option == "y" {
				sd.SaveGame()
			} else {
				fmt.Println("invalid input: aborting")
			}

		case string('r'):
			fmt.Println("roll command sent.")
			fmt.Println("Time to roll")
			if sd.CurrentChapter != core.MAXCHAPTER {
				sd.RollGun()
				sd.PrintData()
			} else {
				fmt.Println("No more chapters, Stranger")
			}
		case string('h'):
			fmt.Println("Printing Help.")
			fmt.Println("r to roll for a gun in the next chapter")
			fmt.Println("l to load existing file")
			fmt.Println("s to save existing data to file")
			fmt.Println("i to print existing game info")
			fmt.Println("q to quit app")
		}
	}
}
