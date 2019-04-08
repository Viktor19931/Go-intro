package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Soccer clubs struct
type SoccerClubs struct {
	Name, Country string
	Value         float32 `json:"Dollar Value(B)"`
	Players       []string
}

var teams = []SoccerClubs{
	{Name: "Machester United", Country: "England", Value: 3.55,
		Players: []string{"David Backham", "Eric antona"}},
	{Name: "Barcelona", Country: "Spain", Value: 7.55,
		Players: []string{"Leo Messi", "Xavi", "Iniesta"}},
	{Name: "Real Madrid", Country: "Spain", Value: 5.55,
		Players: []string{"Zidane", "Ronaldo", "Roberto Carlos"}},
}

func main() {
	fmt.Println("----------------------------------------------------")

	// data, error := json.Marshal(teams)
	data, error := json.MarshalIndent(teams, "", "  ") // better spacing
	fmt.Printf("%s", data)

	if error != nil {
		log.Fatalf("JSON Marshaling failed: %s", error)
	}

	fmt.Println("----------------------------------------------------")
	var names []struct{ Name string }

	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("JSON unmarshall failed: %s", err)
	}

	fmt.Println(names)
}
