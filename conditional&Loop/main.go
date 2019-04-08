package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// ifSwitchPlaying()
	// calculateBmi()

	// checkLeapYear()
	randomRange()
}

func randomRange() {
	const (
		minRandomVal = 10
		maxRandomVal = 22
	)

	var yourRandomVal int
	for {
		fmt.Printf("Enter a random value between %d and %d:", minRandomVal, maxRandomVal)
		fmt.Scanf("%d ", &yourRandomVal)
		if (yourRandomVal < minRandomVal) || (yourRandomVal > maxRandomVal) {
			fmt.Println("*Incorrect input value(s). Please try again ...")
			// continue
		} else {
			break
		}
	}
	rand.Seed(time.Now().UnixNano())
	systemRandomVal := minRandomVal + rand.Intn(maxRandomVal-minRandomVal)

	yourRandomValSumOfDigits := (yourRandomVal / 10) + (yourRandomVal % 10)
	systemRandomValSumOfDigits := (systemRandomVal / 10) + (systemRandomVal % 10)

	lotteryResult := ""

	if systemRandomVal == yourRandomVal {
		lotteryResult = "You won $1000"
	} else if yourRandomValSumOfDigits == systemRandomValSumOfDigits {
		lotteryResult = "You won $500"
	} else {
		lotteryResult = "Try again !"
	}

	fmt.Printf("\nyourRandomVal=%d systemRandomVal=%d \nlotteryResult=%s \n", yourRandomVal, systemRandomVal, lotteryResult)
}

func ifSwitchPlaying() {
	a, b, k := 1, 10, 1
	x := []int{10, 3, 4}
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	for a > b {
		b++
		fmt.Println("for 1")
	}

	// infinit loop
	for {
		if k == 8 {
			break
		}
		if k == 4 || k == 6 {
			k++
			break
		}
		break
	}
	//nested loop
	for i := 1; i < 5; i++ {
		fmt.Printf("\ni=%d ** ", i)
		for j := 1; j < 3; j++ {
			fmt.Printf("%d ", i*j)
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Println("for loop 2")
	}

	for i, val := range x {
		fmt.Println(i, val)
	}

	switch num {
	case 1:
		fmt.Println(num)
	case 2:
		fmt.Println(num)
		fallthrough
	case 3:
		fmt.Println(num)
		fallthrough
	case 4, 5, 6, 7:
		fmt.Println(num)
	default:
		fmt.Println("default")
	}

	switch { // bool tag without statement
	case true:
		fmt.Println("true")
	}

	var c interface{}
	c = 10

	// check by type
	switch i := c.(type) {
	case int:
	case float64:
	case bool, string:
	case func(int) float64:
	case nil:
	default:
		fmt.Println(i)
	}

}

func checkLeapYear() {

	var year int
	fmt.Print("Enter height in year: ")
	fmt.Scanf("%d ", &year)

	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeapYear {
		fmt.Printf("The year %v is a leap \n", year)
	} else {
		fmt.Printf("The year %v is not a leap \n", year)
	}

}

func calculateBmi() {
	var weight float64
	fmt.Print("Enter weight in pounds: ")
	fmt.Scanf("%f ", &weight)

	var height float64
	fmt.Print("Enter height in pounds: ")
	fmt.Scanf("%f ", &height)

	kilogramsPerPound, metersPerInch := 0.45359232, 0.0254

	weightInKilograms := weight * kilogramsPerPound
	heightInMeters := height * metersPerInch

	bmi := weightInKilograms / (heightInMeters * heightInMeters)
	weightStatus := ""

	switch {
	case bmi < 18.5:
		weightStatus = "Underweight"
	case bmi < 25:
		weightStatus = "Normal"
	case bmi < 30:
		weightStatus = "Overweight"
	default:
		weightStatus = "Obese"
	}

	fmt.Printf("%-10s %-10s %-10s %-10s \n", "Weight", "Height", "BMI", "Status")
	fmt.Printf("%-10.2f %-10.2f %-10.2f %-10s \n", weight, height, bmi, weightStatus)
}
