package main

import (
	"fmt"
	"time"
)

func main() {
	scores1 := []float32{91, 82, 99}
	fmt.Printf("average %.2f\n", average(scores1))

	scores2 := []float32{23, 45, 65, 34, 99}
	fmt.Printf("average %.2f\n", average(scores2))

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	names := []string{"Amy", "Jim", "Rob", "Helen"}
	echo(names...)

	data := []string{"Daisy", "Rose", "", "Tulip"}
	fmt.Printf("%q\n", trimSlice(data))

	nums = push(nums, 33)
	fmt.Println(nums)

	nums = pop(nums)
	fmt.Println(nums)

	nums = push(nums, 12, 44)
	fmt.Println(top(nums))

	hi := greeting("ESP")
	fmt.Println(hi())

	hi = greeting("GER")
	fmt.Println(hi())

	add, mult := addCounter(), multBy()

	fmt.Println(add(2))
	fmt.Println(add(-2))
	fmt.Println(mult(3))

	callBacks()
	deferStatment()
	handleError()
}

func handleError() {
	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()

	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)

	panic("Dont panic =)")

}

func deferStatment() {
	fmt.Println("---------------")
	defer fmt.Println("time now ", time.Now())

	fmt.Println("LOC 2", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("LOC 3", time.Now())
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(n, " \n")
	return n * factorial(n-1)
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}

func callBacks() {
	square := func(i int) int {
		return i * i
	}

	cube := func(i int) (res int) {
		res = i * i * i
		return
	}

	fmt.Printf("%v\n", calc(square, 3))
	fmt.Printf("%v\n", calc(cube, 3))

	fmt.Printf("%v\n", calc(func(i int) int {
		return i * i
	}, 3))

}

func calc(f func(int) int, x int) int {
	return f(x)
}

func addCounter() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func multBy() func(int) int {
	total := 1
	return func(i int) (result int) {
		total *= i
		result = total
		return
	}
}

func counter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func greeting(lang string) func() string {
	hi := func() string {
		return "Hello"
	}

	if lang == "ESP" {
		hi = func() string {
			return "Hola"
		}
	} else if lang == "GER" {
		hi = func() string {
			return "Hallo"
		}
	}

	return hi
}

func push(s []int, nums ...int) []int {
	return append(s, nums...)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func top(s []int) int {
	return s[len(s)-1]
}

func trimSlice(data []string) []string {
	var newSlice []string
	for i := range data {
		if data[i] != "" {
			newSlice = append(newSlice, data[i])
		}
	}
	return newSlice
}

func echo(names ...string) {
	for _, s := range names {
		fmt.Print(s, " \n")
	}
}

func average(scores []float32) float32 {
	var total float32

	for _, s := range scores {
		total += s
	}

	return total / float32(len(scores))
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
		fmt.Printf("nums=%v total=%d type=%T\n", nums, total, nums)
	}
}
