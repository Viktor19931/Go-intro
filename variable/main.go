package main

import "fmt"

func main() {
	v, f := 10, 3.14

	fmt.Printf("%v - %T \n", v, v)
	fmt.Printf("%5.3f - %.2f \n ", f, f)

	const pi = 3.14159

	const (
		a = 2
		b
		c
		m = a * 2
	)

	fmt.Println(a, b, c, m)

	const (
		t = 1		//00000001 int 1
		n = t << 1  //00000010 int 2
		s = n << 1  //00000100 int 4
	)

	fmt.Println(t, n, s)

	// %d --> based 10 decimals [0-9]
	// %o --> based 8 octal [0-7]
	// %x --> based 16 hexa decimal [0-9, A-F]
	// %b --> binary [0,1]
	// %T --> represent type of var
	// %v --> value default format
	// %f --> decimal point

	fmt.Printf("%d %[1]o %[1]x %[1]b ", 'A')

	var i uint8 = 20
	var fl float32 = 214.437
	bo := true
	msg := "hello"

	type myString string

	var myMsg myString = "special message"

	fmt.Printf("%d %[1]v %[1]T \n", i)
	fmt.Printf("%5.1f %.4[1]f %[1]g %[1]T \n", fl)
	fmt.Printf("%t %[1]v %[1]T \n", bo)
	fmt.Printf("%s %[1]T \n", msg)
	fmt.Printf("%v %[1]T \n", myMsg)

}
