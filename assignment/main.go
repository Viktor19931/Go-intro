package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	num1, num2, num3 := 12, 4, 5
	fahrenheit := 32.1
	radius := 12.54
	i := 2

	fmt.Println("Average of ", num1, num2, num3, "is ", (num1+num2+num3)/3)
	fmt.Printf("fahrenheit is %5.2f and celsius is %5.4f\n", fahrenheit, fahrenheitToCelsius(fahrenheit))
	fmt.Printf("The area for the circle of radius %5.2f is %5.2f\n", radius, calculateCircleArea(radius))

	displayTable(i)
	// checkDay()
	// checkTime()
	// random()
	printRune()
}

func printRune() {
	s1, s2 := "漢字", "漢字"
	var myRunes []rune

	for _, r := range s1 {
		myRunes = append(myRunes, r)
	}

	for _, r := range s2 {
		myRunes = append(myRunes, r)
	}

	fmt.Printf("%q\n", myRunes)
	fmt.Printf("%q\n", string(myRunes))
}

func random() {
	const (
		a = iota
		b = 5
		c = iota
		d = 8
		e = iota
	)

	fmt.Printf("a=%d b=%d c=%d d=%d e=%d\n", a, b, c, d, e)
}

func checkTime() {
	totalMilisec := time.Now().UnixNano() / int64(time.Millisecond)

	totalSec := totalMilisec / 1000
	currentSeconds := totalSec % 60
	totalMins := totalSec / 60
	currentMins := totalMins % 60
	totalHours := totalMins / 60
	currentHours := totalHours % 24

	fmt.Printf("tSec = %d\n", totalSec)
	fmt.Printf("cSec = %d\n", currentSeconds)
	fmt.Printf("tMins = %d\n", totalMins)
	fmt.Printf("cMins = %d\n", currentMins)
	fmt.Printf("tHours = %d\n", totalHours)
	fmt.Printf("cHours = %d\n", currentHours)
}

func checkDay() {
	today, duration := 0, 7
	remainingDays := (today + duration) % 7
	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Suturday"}
	fmt.Println("If Today Sunday, after 85 days will be\n", days[remainingDays])
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func calculateCircleArea(r float64) float64 {
	const PI = math.Pi
	return PI * math.Pow(r, 2)
}

func displayTable(i int) {
	fmt.Printf("i \t i*2 \t i*3\n")
	fmt.Printf("==\t ===\t ===\t\n")

	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
}
