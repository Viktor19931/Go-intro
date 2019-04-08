package myutil

// AAA
func Area(l, w uint) uint {
	return l * w
}

// BBB
func Perim(l, w uint) uint {
	return 2*(l + w)
}

// AAA
func Average(numbers []float32) float32 {
	total :=float32(0)
	for _, n := range numbers {
		total += n
	}
	return total / float32(len(numbers))
}