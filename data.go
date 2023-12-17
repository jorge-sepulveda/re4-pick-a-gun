package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

const MAXCHAPTER = 16
const STARTCHAPTER = 1

var handguns = []string{"SR-09 R", "Punisher", "Red9", "Blacktail", "Matilda", "Sentinel Nine"}
var shotguns = []string{"W-870", "Riot Gun", "Striker", "Skull Shaker"}
var rifles = []string{"SR M1903", "Stingray", "CQBR Assault Rifle"}
var magnums = []string{"Broken Butterfly", "Killer7"}

//var specials = []string{"Bolt Thrower", "Infinite Rocket Launcher", "Chicago Sweeper"}

// SaveData
type SaveData struct {
	CurrentChapter int      `json:"current_chapter"`
	CurrentGun     string   `json:"current_gun"`
	UsedGuns       []string `json:"used_guns"`
	GunsList       []string `json:"guns_list"`
}

// StartGame initializes and starts the game with a selected gun
func (s *SaveData) StartGame() error {
	s.CurrentChapter = STARTCHAPTER
	s.GunsList = append(s.GunsList, handguns...)
	s.GunsList = append(s.GunsList, shotguns...)
	s.GunsList = append(s.GunsList, rifles...)
	s.GunsList = append(s.GunsList, magnums...)
	s.GunsList = append(s.GunsList, handguns...)
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
	data, _ := json.MarshalIndent(s, "", "\t")
	err := os.WriteFile("data.json", data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully saved to data.json\n")
	return nil
}

func (s *SaveData) LoadGame() error {
	file, _ := os.ReadFile("data.json")
	err := json.Unmarshal([]byte(file), &s)
	if err != nil {
		return err
	}
	return nil
}
