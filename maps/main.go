package main

import (
	"fmt"
	"sort"
)

var players = make(map[string]map[string]bool)

func main() {
	days := make(map[string]int)
	fmt.Println(days)

	colors := map[string]string{
		"red":   "ff0000",
		"green": "#4b745",
	}

	employees := map[string]float64{
		"Blacke": 60000.00,
		"Parker": 12000.50,
		"Dakota": 93000.00,
	}
	// sort map
	names := make([]string, 0, len(employees))
	for n := range employees {
		names = append(names, n)
	}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, employees[n])
	}

	salary1, ok := employees["Parker"] // ok is a failure or success look up

	if ok {
		fmt.Println("success to look up", salary1)
	} else {
		fmt.Println("faile to look up")
	}

	_, ok = employees["Dakota"]
	if ok {
		// do something
	}

	for name, salary := range employees {
		fmt.Println(name, salary)
	}

	colors["white"] = "#fff"

	delete(colors, "white")

	printMap(colors)
	generateMap()

	mapOfMaps()

	addPlayer("Leo", "Messi")
	addPlayer("Roger", "Federer")
	addPlayer("Michael", "Jordan")
	fmt.Println(players)

	
	fmt.Println(hasPlayer("Leo", "Messi"))
	fmt.Println(hasPlayer("fff", "Messi"))
}

func addPlayer(fn, ln string) {
	n := players[fn]
	if n == nil {
		n = make(map[string]bool)
		players[fn] = n
	}
	n[ln] = true
}

func hasPlayer(fn, ln string) bool {
	return players[fn][ln]
}

func mapOfMaps() {
	emloyees := map[string]map[string]string{
		"BT": map[string]string{
			"firstName": "Parker",
			"lastName":  "Cooper",
		},
		"PC": map[string]string{
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"DC": map[string]string{
			"firstName": "Dakota",
			"lastName":  "Carrington",
		},
	}

	if emp, ok := emloyees["PC"]; ok {
		fmt.Println(emp["firstName"])
	}

	for initials, emp := range emloyees {
		fmt.Println(initials, emp["firdtName"], emp["lastName"])
	}
}

func checkDublicate() {
	str := []string{"One", "Two", "Two", "Three", "Three", "Five"}

	m := make(map[string]bool)

	for i := range str {
		word := str[i]
		if !m[word] {
			m[word] = true
		}
	}
}

func generateMap() {
	pow2 := make(map[int]int)

	for i := 1; i < 10; i++ {
		pow2[i] = i * i
	}

	fmt.Println(pow2, len(pow2))
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "-->", hex)
	}
}
