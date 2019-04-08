package main

import (
	"fmt"
	an "packages/animals"
	"packages/athletes"
	"packages/shapes"
)

func init() {}

func main() {
	testCase3()
}

func testCase1() {
	fmt.Printf("area --> %d\n", shapes.Area(2, 4))
	fmt.Printf("shapes length=%d, width=%d\n", shapes.Length, shapes.Width)
	fmt.Printf("square area %d\n", shapes.SquareArea(4))
}

func testCase2() {
	human := an.Human{}
	fmt.Println(human.Speaks())

	lion := an.Lion{}
	fmt.Println(lion.Speaks())

	cat := an.Cat{}
	fmt.Println(cat.Speaks())
}

func testCase3() {
	player1 := athletes.Player{"Name", "Sport", 22, athletes.Info{"Country", "HairColor"}}
	fmt.Println(player1)

	
}
