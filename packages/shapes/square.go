package shapes

import "fmt"

// SquareArea calculation
func SquareArea(s int) int {
	return s * s
}

func init() {
	fmt.Println("=> init() from shapes")
}