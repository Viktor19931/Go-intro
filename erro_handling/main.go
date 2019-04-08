package main

import (
	"fmt"
	"log"
	"os"
)

var fileName = "non_exist.txt"

func main() {
	testCase3()
	fmt.Println("control has returned to the main function")
}

func testCase1() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Println("Error: ", err)
	}
}

func testCase2() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
}
}

func testCase3() {
	_, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
}