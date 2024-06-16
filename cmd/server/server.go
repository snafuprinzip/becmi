package main

import (
	"becmi"
	"becmi/dice"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Monster struct {
	ID     uint     `json:"id"`
	Name   string   `json:"name"`
	Powers []string `json:"powers"`
}

func monsters() []Monster {
	return []Monster{
		{
			ID:     1,
			Name:   "Dracula",
			Powers: []string{"Immortality", "Shape-shifting", "Mind Control"},
		},
		{
			ID:     2,
			Name:   "Frankenstein",
			Powers: []string{"Superhuman Strength", "Endurance"},
		},
		{
			ID:     3,
			Name:   "Werewolf",
			Powers: []string{"Shape-shifting", "Enhanced Senses", "Regeneration"},
		},
		{
			ID:     4,
			Name:   "Zombie",
			Powers: []string{"Undead Physiology", "Immunity to Pain"},
		},
		{
			ID:     5,
			Name:   "Mummy",
			Powers: []string{"Immortality", "Control over Sand"},
		},
	}
}

func loadMonsters() map[string]Monster {
	monsters := monsters()
	res := make(map[string]Monster, len(monsters))

	for _, x := range monsters {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}

var characters map[string]becmi.Character

func RootGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func CharacterGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handling READ request - Method:", r.Method)
	monster, exists := loadMonsters()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//json.NewEncoder(w).Encode(monster)
	b, _ := json.MarshalIndent(monster, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

	//characterID := r.PathValue("id")
	//character, ok := monsters[characterID]
	//if ok {
	//	b, _ := json.MarshalIndent(character, "", "  ")
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write(b)
	//} else {
	//	http.NotFound(w,r)
	//}
}

func main() {
	dice.Roll("3w6-1")
	//mux := http.NewServeMux()
	//mux.HandleFunc("GET /", RootGetHandler)
	//mux.HandleFunc("GET /character/{id}", CharacterGetHandler)
	//
	//fmt.Println("Listening on port 8080")
	//log.Fatal(http.ListenAndServe(":8080", mux))
}
