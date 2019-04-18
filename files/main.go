package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type team []string
type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

var fileName = "names.txt"
var newFileName = "new.txt"

func main() {
	// writeToFile()
	// readFromFile()
	// renameFile()
	// removeFile()

	// writeToJson()
	// readFile()
	readWords()
}

func readFile() {
	// go run main.go test.json arg2 arg3  --> example of run
	f, err := os.Open(os.Args[1])

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	fmt.Println(f)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, f)
	}

	currentDirectory()
}

func currentDirectory() {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println("\nCurrent directory=", cwd)
}

func readWords() {
	input := bufio.NewScanner(os.Stdin)

	var all string
	fmt.Println("Enter some words. To exist, type 'end'.")

	for input.Scan() {
		s := input.Text()

		if s == "end" {
			break
		}

		all += (s + " ")
	}

	fmt.Println(all)
}

func writeToJson() {
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0666)
}

func writeToFile() {
	players := team{"Messi", "Pele", "Maradonna", "Zidane", "Platini"}

	err := ioutil.WriteFile(fileName, []byte(players.toString()), 0666)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Wrote the following to file: %v \n", players.toString())
	}
}

func (t team) toString() string {
	return strings.Join([]string(t), ",")
}

func readFromFile() {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	s := strings.Split(string(bs), ",")
	fmt.Println(s)
}

func renameFile() {
	err := os.Rename(fileName, newFileName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File renamed !")
	}
}

func removeFile() {
	err := os.Remove(newFileName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File renamed !")
	}
}
