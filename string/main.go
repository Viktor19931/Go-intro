package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello World !"
	ch := "Hello 漢字"

	fmt.Println(len(ch))                    // reprecent bytes count
	fmt.Println(utf8.RuneCountInString(ch)) // represent chars count
	fmt.Println([]rune(ch))
	fmt.Println(string([]rune(ch)))

	fmt.Println(len(str))
	fmt.Println(str[5]) // represent 5th chatacter

	fmt.Println(str[:5])
	fmt.Println(str[5:])
	fmt.Println(str[2:9])

	fmt.Println(str[:6] + str[len(str)-1:])

	// compare()
	testCase()
}

func testCase() {
	fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   "))

	fmt.Printf("Fields are: %q \n", bytes.Fields([]byte("  foo bar  baz   ")))

	fmt.Println(strings.HasSuffix("friend", "end"))
	fmt.Println(strings.HasPrefix("Golang", "Go"))

	fmt.Println(strings.Index("booklet", "let"))
	fmt.Println(strings.Index("booklet", "lets"))

	s := []string{"Mon", "Tue", "Wed"}

	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Index("go leafs go", "go"))
	fmt.Println(strings.LastIndex("go leafs go", "go"))

	shiftOne := func(r rune) rune{
		return r + 1
	}

	fmt.Println(strings.Map(shiftOne, "ABCDabcd I am going to be converter"))
	fmt.Println(strings.Repeat("la", 3))
	
	fmt.Println(strings.Replace("blah blah blah", "h", "ck", 2)) // -1 for unlimited replacement
	fmt.Println(strings.Split("a,b,c", ","))
}

func compare() {
	fmt.Println(strings.Compare("at", "at"))
	fmt.Println(strings.Compare("at", "amber"))
	fmt.Println(strings.Compare("amber", "at"))

	fmt.Println(bytes.Compare([]byte("team player"), []byte("play")))

	fmt.Println(strings.Count("render", "e"))
}
