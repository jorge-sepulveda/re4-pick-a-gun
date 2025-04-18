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
	var ready = false

	fmt.Println("Welcome to the RE4 Pick a gun service!")
	fmt.Println("Press l or a to start Leon's or Ada's playthrough or h for help")

	for {
		fmt.Scan(&option)
		fmt.Println("Choose your option, Stranger: ")
		switch option {
		case string('l'):
			fmt.Println("Starting leon's playthrough...")
			err := sd.StartGame("L", core.Handguns, core.Shotguns, core.Rifles, core.Subs, core.Magnums)
			if err != nil {
				print(err.Error() + "\n")
				os.Exit(1)
			}
			ready = true
		case string('a'):
			fmt.Println("Starting Ada's playthrough...")
			err := sd.StartGame("A", core.AdaHandguns, core.AdaShotguns, core.AdaRifles, core.AdaSubs, core.AdaSpecials)
			if err != nil {
				print(err.Error() + "\n")
				os.Exit(1)
			}
			ready = true
		case string('q'):
			fmt.Println("Quit command sent.")
			os.Exit(0)
		case string('h'):
			fmt.Println("Printing Help.")
			fmt.Println("a to roll for Seperate Ways")
			fmt.Println("l to roll Main playthrough")
			fmt.Println("q to quit app")
		}
		if ready {
			break
		}
	}

	fmt.Printf("Starting in Chapter %d\n", sd.CurrentChapter)
	fmt.Printf("Current Gun: %s\n", sd.CurrentGun)
	fmt.Printf("Used guns list %+v\n", sd.UsedGuns)
	fmt.Printf("existing guns list %v\n", sd.GunsList)
	for {
		fmt.Println("Choose your weapon, Stranger: (h for help)")
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
			err := sd.LoadGame()
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
			if sd.CurrentChapter != sd.FinalChapter {
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
