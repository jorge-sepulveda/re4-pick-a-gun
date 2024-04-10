// Paxkage core provides the main functionality. Loading, saving and rolloing guns  are inside this.
package core

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const STARTCHAPTER = 1
const MAXCHAPTER = 16

var (
	Handguns = []string{"SR-09 R", "Punisher", "Red9", "Blacktail", "Matilda", "Sentinel Nine"}
	Shotguns = []string{"W-870", "Riot Gun", "Striker", "Skull Shaker"}
	Rifles   = []string{"SR M1903", "Stingray", "CQBR Assault Rifle"}
	Magnums  = []string{"Broken Butterfly", "Killer7"}
	Subs     = []string{"TMP", "LE 5"}
	//ass = [string]{"Bolt Thrower"}
	Specials = []string{"Handcannon", "Infinite Rocket Launcher", "Chicago Sweeper"}
)

// SaveData
type SaveData struct {
	CurrentChapter int      `json:"current_chapter"`
	CurrentGun     string   `json:"current_gun"`
	UsedGuns       []string `json:"used_guns"`
	GunsList       []string `json:"guns_list"`
	Checks         CheckedGuns
}

type CheckedGuns struct {
	handguns   bool
	shotguns   bool
	rifles     bool
	magnums    bool
	subs       bool
	specials   bool
	deletePool bool
}

// StartGame initializes the gun pool, randomizing and starting the game with a selected gun
func (s *SaveData) StartGame(guns ...[]string) error {
	s.CurrentChapter = STARTCHAPTER
	for i := range guns {
		s.GunsList = append(s.GunsList, guns[i]...)
	}
	//error check in the event there aren't enough guns for all the chapters.
	if len(s.GunsList) < MAXCHAPTER {
		return fmt.Errorf("ERROR: Not enough guns in the pool")
	}
	rand.Shuffle(len(s.GunsList), func(i, j int) {
		s.GunsList[i], s.GunsList[j] = s.GunsList[j], s.GunsList[i]
	})
	s.GunsList = s.PickGun()
	return nil
}

// Rollgame increments the current chapter and randomly selects a weapon from the pool.
func (s *SaveData) RollGun() error {
	s.CurrentChapter += 1
	s.GunsList = s.PickGun()
	return nil
}

// PickGun randomly picks a weapon from the pool and returns the array with the weapon removed.
// Internal helper function
func (s *SaveData) PickGun() []string {
	ridx := rand.Intn(len(s.GunsList))
	s.CurrentGun = s.GunsList[ridx]
	s.UsedGuns = append(s.UsedGuns, s.CurrentGun)
	return append(s.GunsList[:ridx], s.GunsList[ridx+1:]...)
}

// SaveGame marshalls the existing savedata object
func (s *SaveData) SaveGame() error {
	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully saved to data.json\n")
	return nil
}

// LoadGame loads the current saved data into the app,
func (s *SaveData) LoadGame() error {
	backup, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("ERROR: Could not backup data: %s", err)
	}
	file, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("ERROR: Could not read file: %s", err)
	}
	err = json.Unmarshal([]byte(file), &s)
	if err != nil {
		return fmt.Errorf("ERROR: could not load file: %s", err)
	}
	if (MAXCHAPTER - s.CurrentChapter) > len(s.GunsList) {
		err = json.Unmarshal(backup, &s)
		if err != nil {
			return fmt.Errorf("ERROR: Could not reload backup: %s", err)
		}
		return fmt.Errorf("ERROR: not enough weapons in gunpool. Reverting")
	}
	return nil
}

// PrintData prints the existing struct details.
func (s *SaveData) PrintData() {
	fmt.Printf("Chapter: %d\n", s.CurrentChapter)
	fmt.Printf("Selected gun: %s\n", s.CurrentGun)
	fmt.Printf("Used guns list: %v\n", s.UsedGuns)
	//fmt.Printf("Existing guns list: %v\n", s.GunsList)
}

func (s *SaveData) SetHandguns(b bool) {
	s.Checks.handguns = b
	log.Printf("handguns set to %t\n", s.Checks.handguns)
}

func (s *SaveData) SetShotguns(b bool) {
	s.Checks.shotguns = b
}

func (s *SaveData) SetRifles(b bool) {
	s.Checks.rifles = b
}

func (s *SaveData) SetMagnums(b bool) {
	s.Checks.magnums = b
}

func (s *SaveData) SetSubs(b bool) {
	s.Checks.subs = b
}

func (s *SaveData) SetSpecials(b bool) {
	s.Checks.specials = b
}

func (s *SaveData) GetHandguns() bool {
	return s.Checks.handguns
}

func (s *SaveData) GetShotguns() bool {
	return s.Checks.shotguns
}

func (s *SaveData) GetRifles() bool {
	return s.Checks.rifles
}

func (s *SaveData) GetMagnums() bool {
	return s.Checks.magnums
}

func (s *SaveData) GetSubs() bool {
	return s.Checks.subs
}

func (s *SaveData) GetSpecials() bool {
	return s.Checks.specials
}
