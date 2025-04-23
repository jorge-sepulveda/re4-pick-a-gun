package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/start", PingHandler).Methods("GET")
	r.HandleFunc("/roll", rollHandler).Methods("POST")
	r.HandleFunc("/load", loadHandler).Methods("POST")
	http.ListenAndServe(":8080", r)

}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}

	sd.StartGame("L", core.Handguns, core.Shotguns, core.Rifles, core.Subs, core.Magnums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}

func rollHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}
	err := json.NewDecoder(r.Body).Decode(&sd)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if sd.CurrentChapter == sd.FinalChapter {
		http.Error(w, "All out of chapters, stranger!", http.StatusOK)
		return

	}
	sd.RollGun()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Payload Data, unable to load", http.StatusBadRequest)
	}
	defer r.Body.Close()
	sd.LoadString(requestBytes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}
