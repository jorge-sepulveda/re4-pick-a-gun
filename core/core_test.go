package core

import (
	"log"
	"testing"
)

func TestStartGame(t *testing.T) {
	testSD := SaveData{}
	err := testSD.StartGame()
	if err == nil {
		log.Println("Expected Empty gun array error")
		t.Fail()
	}
	testSD.SetHandguns(true)
	err = testSD.StartGame()
	if testSD.DisableDelete == false {
		log.Println("Too few guns in the pool, disable delete must be enabled")
		t.Fail()
	}
	if len(testSD.GunsList) == 0 {
		log.Println("Gun list should be greater than size 1")
		t.Fail()
	}

}

func TestPickGun(t *testing.T) {
	testSD := SaveData{
		GunsList: Handguns,
	}
	newGunsList := testSD.PickGun()
	if len(newGunsList) != len(Handguns)-1 {
		log.Println("Gun array is mismatching sizes")
		t.Fail()
	}
	if testSD.CurrentGun == "" {
		log.Println("Selected gun is nil")
		t.Fail()
	}
}

func TestRollGun(t *testing.T) {
	testSD := SaveData{}
	testSD.CurrentChapter = 87
	err := testSD.RollGun()
	if err == nil {
		log.Println("Chapter exceeds max chapter")
	}
	testSD.CurrentChapter = 1
	err = testSD.RollGun()
	if err == nil {
		log.Println("Expected empty gun list error")
	}
	testSD.GunsList = Handguns
	err = testSD.RollGun()
	gunLength := len(testSD.GunsList)
	if err != nil || testSD.CurrentChapter != 2 || gunLength != gunLength-1 {
		log.Println("Can't do math nor pick a gun")
	}

}

func TestPrintData(t *testing.T) {
	testSD := SaveData{}
	testSD.PrintData()
}

func TestBuildGunPool(t *testing.T) {
	testSD := SaveData{}
	testSD.SetHandguns(true)
	testSD.SetShotguns(true)
	testSD.SetRifles(true)
	testSD.SetMagnums(true)
	testSD.SetSubs(true)
	testSD.SetSpecials(true)
	gunPool := testSD.BuildGunPool()
	if len(gunPool) != 20 {
		log.Println("Failed to build complete pool")
		t.Fail()
	}
}
