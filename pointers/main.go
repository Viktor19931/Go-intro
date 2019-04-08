package main

import "fmt"

func main() {
	// * --> mean content
	// & --> mean address

	x := 1
	p := &x

	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T\n", x, &x, p, *p, &p)

	c := []int{2,4,6,8,10}
	b := &c

	fmt.Printf("x=%v *p=%d &x=???=%x p???=%x &p=%x\n", c, *b, &c, b, &c)

	s := [...]string{"Angela", "Tim", "Nicole", "Henry"}
	k := &s

	y := s[1:3]
	q := &y

	fmt.Printf("*p=%s &p=%x\n", *k, &k)
	fmt.Printf("*p=%s &p=%x\n", *q, &q)

	(*k)[1] = "NICOLE"
	(*q)[1] = "TIM"

	fmt.Println(s)
	fmt.Println(y)

	pointerSlices()

	t := 10

	fmt.Printf("t=%d\n", t)
	f2(&t)
	fmt.Printf("t=%d\n", t)

	assignment()
}

func assignment(){
	a := 1
	fmt.Printf("(1) a=%d &a=%x \n\n", a, &a )

	f3(&a)
	fmt.Printf("(2) a=%d \n\n", a)
	fmt.Printf("(3) a=%d\n", f3(&a))
}

func f3(y *int) int{
	*y++
	fmt.Printf("(f) *y=%d y=%x\n", *y, y)
	return *y
}

func f2(y *int) {
	fmt.Printf("(f2) y=%d\n", *y)
	*y += 2
	fmt.Printf("(f2) y=%d\n", *y)
}

func pointerSlices() {
	slices := [][]int{{2,3,4},{5,6},{7,8,9,10}}
	var bigSlice []int

	var p *[][]int
	p = &slices
	for index := range *p {
		bigSlice = append(bigSlice, (*p)[index]...)
	}

	fmt.Printf("%v %T\n", *p, *p)
	fmt.Printf("%v %T\n", bigSlice, bigSlice)

	var p1 = f(2)
	fmt.Println(p1, *p1, &p1)

	var p2 = f(3)
	fmt.Println(p2, *p2, &p2)
}

func f(inp int) *int{
	v := inp * 2
	fmt.Printf("&v=%x\n", &v)

	return &v
}