package main

import "fmt"

type rectangle struct {
	w, l int
}

type square struct {
	s int
}

type shape interface {
	area() int
	perim() int
}

type athlete struct {
	name, country string
}

type football struct {
	athlete
	position string
}

type tennis struct {
	athlete
	rightHanded bool
}

type athletes interface {}

func main() {
	r1 := rectangle{2,4}
	r2 := rectangle{4,6}

	s1 := square{3}
	s2 := square{6}

	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

	fmt.Printf("Total area =%d\n", totalArea(&r1, &r2, &s1, &s2))

	emptyExample()
}

func emptyExample() {

	var num interface{} = 10
	fmt.Println(num.(int) + 12)
	
	messi := football{}
	pele := football{}
	federar := tennis{}
	nadal := tennis{}

	favAthletes := []athletes{messi, pele, federar, nadal}

	for k, v := range favAthletes {
		fmt.Println(k, " - ", v)
	}

	messi = football{athlete{"Leo Messi", "Argentina"}, "Attacker"}
	federar = tennis{athlete{"Roger Federar", "Switzerland"}, true}
	playerInfo(messi)
	playerInfo(federar)

	pele = football{athlete{"Pele", "Argentina"}, "Attacker"}
	nadal = tennis{athlete{"Nadal", "Switzerland"}, true}

	favAthletes2 := []interface{}{messi, pele, federar, nadal}

	fmt.Println(favAthletes2)
}

func playerInfo(a interface{}) {
	fmt.Println(a)
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.w + c.l)
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

func info(s shape) {
	fmt.Printf("area()=%d perim()=%d\n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}