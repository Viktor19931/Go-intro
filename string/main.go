package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello World !"
	ch := "Hello 漢字"

	fmt.Println(len(ch)) // reprecent bytes count
	fmt.Println(utf8.RuneCountInString(ch)) // represent chars count
	fmt.Println([]rune(ch))
	fmt.Println(string([]rune(ch)))

	fmt.Println(len(str))
	fmt.Println(str[5]) // represent 5th chatacter

	fmt.Println(str[:5])
	fmt.Println(str[5:])
	fmt.Println(str[2:9])

	fmt.Println(str[:6] + str[len(str)-1:])

}
