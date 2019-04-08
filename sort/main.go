package main

import (
	"fmt"
	"sort"
)

func main() {
	sortInts()
	sortStrings()
	ex1()
}

func ex1() {
	
}

func sortInts() {
	n := []int{2, 7, 4, 7, 1, 0, 5, 12}

	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

func sortStrings() {
	s := []string{"ccc", "ddd", "mmm", "aaa"}
	fmt.Printf("\n%s \n", s)
	sort.Strings(s)
	fmt.Printf("\n%s \n", s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
