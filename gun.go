package main

import (
	"fmt"
)

func main() {
	var sd SaveData
	var option string
	sd.StartGame()
	fmt.Println("Welcome to the RE4 Pick a gun service!")
	fmt.Printf("Current Gun: %s\n", sd.CurrentGun)
	fmt.Printf("Used guns list %v\n", sd.UsedGuns)
	fmt.Printf("existing guns list %v\n", sd.GunsList)
	for {
		fmt.Println("Choose your weapon, Stranger: ")
		fmt.Scan(&option)
		switch option {
		case string('q'):
			fmt.Println("Quit command sent.")
			return
		case string('l'):
			fmt.Println("Load command sent.")
			sd.LoadGame()
			fmt.Printf("Starting Chapter: %d\n", sd.CurrentChapter)
			fmt.Printf("Selected gun: %s\n", sd.CurrentGun)
			fmt.Printf("Used guns list: %v\n", sd.UsedGuns)
			fmt.Printf("Existing guns list: %v\n", sd.GunsList)
		case string('s'):
			fmt.Println("Save command sent.")
			sd.SaveGame()
		case string('r'):
			fmt.Println("roll command sent.")
			if sd.CurrentChapter != MAXCHAPTER {
				sd.RollGun()
				fmt.Printf("Starting Chapter: %d\n", sd.CurrentChapter)
				fmt.Printf("Selected gun: %s\n", sd.CurrentGun)
				fmt.Printf("Used guns list: %v\n", sd.UsedGuns)
				fmt.Printf("Existing guns list: %v\n", sd.GunsList)
			} else {
				fmt.Println("No More chapters, Stranger")
			}

		case string('h'):
			fmt.Println("Printing Help.")
			fmt.Println("r to roll for a gun in the next chapter")
			fmt.Println("l to load existing file")
			fmt.Println("s to save existing data to file")
			fmt.Println("q to quit app")

		}
	}
}
