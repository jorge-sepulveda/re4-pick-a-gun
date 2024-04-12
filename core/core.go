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
	DisableDelete  bool     `json:"disable_delete"`
	Checks         CheckedGuns
}

type CheckedGuns struct {
	Handguns bool `json:"handguns"`
	Shotguns bool `json:"shotguns"`
	Rifles   bool `json:"rifles"`
	Magnums  bool `json:"magnums"`
	Subs     bool `json:"subs"`
	Specials bool `json:"specials"`
}

// StartGame initializes the gun pool, randomizing and starting the game with a selected gun
func (s *SaveData) StartGame() error {
	s.CurrentChapter = STARTCHAPTER
	s.GunsList = s.BuildGunPool()
	//error check in the event there aren't enough guns for all the chapters.
	if len(s.GunsList) < MAXCHAPTER {
		fmt.Printf("%v\n", s.GunsList)
		fmt.Printf("Not enough guns in the pool. Disabling deleting from pool\n")
		s.DisableDelete = true
	}
	fmt.Println(len(s.GunsList))
	if len(s.GunsList) == 0 {
		return fmt.Errorf("No guns in the pool, select at least one option")
	}
	rand.Shuffle(len(s.GunsList), func(i, j int) {
		s.GunsList[i], s.GunsList[j] = s.GunsList[j], s.GunsList[i]
	})
	s.GunsList = s.PickGun()
	fmt.Printf("%+v\n", s)
	return nil
}

// Rollgame increments the current chapter and randomly selects a weapon from the pool.
func (s *SaveData) RollGun() error {
	if s.CurrentChapter == MAXCHAPTER {
		return fmt.Errorf("All out of chapters, Stranger!")
	}
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
	if s.DisableDelete {
		return s.GunsList
	}
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
	if (MAXCHAPTER-s.CurrentChapter) > len(s.GunsList) && !s.DisableDelete {
		err = json.Unmarshal(backup, &s)
		if err != nil {
			return fmt.Errorf("ERROR: Could not reload backup: %s", err)
		}
		return fmt.Errorf("ERROR: not enough weapons in gunpool. Reverting")
	}
	fmt.Printf("%+v\n", s)
	return nil
}

// PrintData prints the existing struct details.
func (s *SaveData) PrintData() {
	fmt.Printf("Chapter: %d\n", s.CurrentChapter)
	fmt.Printf("Selected gun: %s\n", s.CurrentGun)
	fmt.Printf("Used guns list: %v\n", s.UsedGuns)
	//fmt.Printf("Existing guns list: %v\n", s.GunsList)
}

func (s *SaveData) BuildGunPool() []string {
	gunPool := []string{}
	if s.GetHandguns() {
		gunPool = append(gunPool, Handguns...)
	}
	if s.GetShotguns() {
		gunPool = append(gunPool, Shotguns...)
	}
	if s.GetRifles() {
		gunPool = append(gunPool, Rifles...)
	}
	if s.GetMagnums() {
		gunPool = append(gunPool, Magnums...)
	}
	if s.GetSubs() {
		gunPool = append(gunPool, Subs...)
	}
	if s.GetSpecials() {
		gunPool = append(gunPool, Specials...)
	}
	return gunPool
}

func (s *SaveData) SetHandguns(b bool) {
	s.Checks.Handguns = b
	log.Printf("handguns set to %t\n", s.Checks.Handguns)
}

func (s *SaveData) SetShotguns(b bool) {
	s.Checks.Shotguns = b
}

func (s *SaveData) SetRifles(b bool) {
	s.Checks.Rifles = b
}

func (s *SaveData) SetMagnums(b bool) {
	s.Checks.Magnums = b
}

func (s *SaveData) SetSubs(b bool) {
	s.Checks.Subs = b
}

func (s *SaveData) SetSpecials(b bool) {
	s.Checks.Specials = b
}

func (s *SaveData) GetHandguns() bool {
	return s.Checks.Handguns
}

func (s *SaveData) GetShotguns() bool {
	return s.Checks.Shotguns
}

func (s *SaveData) GetRifles() bool {
	return s.Checks.Rifles
}

func (s *SaveData) GetMagnums() bool {
	return s.Checks.Magnums
}

func (s *SaveData) GetSubs() bool {
	return s.Checks.Subs
}

func (s *SaveData) GetSpecials() bool {
	return s.Checks.Specials
}
