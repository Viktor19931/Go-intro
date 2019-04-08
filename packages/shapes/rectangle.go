package shapes

import (
	"fmt"
	"log"
)

// Length
var Length = 2

// Width
var Width = 3

// Area func
func Area(l, w int) int {
	return l * w
}

func init() {
	fmt.Println("init in rectangle")

	if (Length < 0) || (Width < 0) {
		log.Fatal("Length & Width can not be less than zero")
	}
}
