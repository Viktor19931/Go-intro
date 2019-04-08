package main

import (
	"fmt"
	"strings"
	"time"

	athletes "./types"
	"github.com/fxtlabs/date"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type generalIfo struct {
	country, hairColor string
}

type player struct {
	name, sport string
	age         int
}

type employee struct {
	name, job    string
	lastLoggedIn string
	DOB          date.Date
}

type book struct {
	name   string
	author string
}

type hunter struct {
	person
	stumina, power string
}

type movie struct {
	name, actor string
}

type imdb struct {
	movie
	name, comment string
}

type artist struct {
	name string
	age  int
}

type singer struct {
	field string
	artist
}

func main() {
	// initialization
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "email@com.ua",
			zipCode: 23132,
		},
	}

	// var alex person

	// alex.firstName = "Vasd"
	// alex.lastName = "Fgerg"
	alex.updateName("Viktor")
	alex.print()

	example1()
	example2()

	b1 := book{"fdsdsfdsf", "fdsfdsdsf"}
	b2 := book{"fdsdsfdsf", "fdsfdsdsf"}
	b3 := book{"fdsdsfdfdfsf", "fdsffdfdsdsf"}

	fmt.Println(compare1(b1, b2))
	fmt.Println(compare1(b1, b3))

	fmt.Println(compare2(&b1, &b2))

	assignment()
	player1 := athletes.Player{"Anderson Silva", "MMA", 43, athletes.Info{"Brazil", "Black"}}
	fmt.Println("(1) player1:", player1)

	changeAthleteName2(&player1)
	fmt.Println("(2) player1:", player1)

	changeAthleteName1(player1)
	fmt.Println("(3) player2:", *player1.ToLowercase())

	lastExample()
}

func lastExample() {
	r1 := artist{
		"Michael Jackson",
		32,
	}

	fmt.Println(r1)
	r1.showName()
	r1.showAge()

	s1 := singer{
		field: "singer",
		artist: r1,
	}

	fmt.Printf("\n%v\n", s1)
	s1.showName()
	s1.showAge()
	s1.showfield()
}

func (a *artist) showName() {
	fmt.Println("Name=", a.name)
}

func (a *artist) showAge() {
	fmt.Println("Age=", a.age)
}

func (a *singer) showAge() {
	fmt.Println("Fake Age=", a.age+100)
}

func (a *singer) showfield() {
	fmt.Println("Field=", a.field)
}

func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func (i imdb) fullInfo() string {
	return strings.ToLower(i.movie.name + "-" +
		i.movie.actor + "-" + i.name + "-" + i.comment)
}

func changeAthleteName1(p athletes.Player) {
	p.Name = "Anderson Silva"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "Viktor Berny"
	p.Country = strings.ToUpper(p.Country)
}

func assignment() {
	var a = new(hunter)
	a.firstName = "Viktor"
	a.lastName = "Hardubej"
	a.zipCode = 111111
	a.email = "email@yahoo.com"
	a.stumina = "70%"
	a.power = "hight"

	fmt.Printf("value=%#v\n", a) // <-- nice form of printing ########

	m1 := movie{name: "Forest Gump", actor: "Tom Hanks"}
	m2 := movie{name: "The Godfather", actor: "Marlon Brando"}

	fmt.Println(m1.fullInfo())
	fmt.Println(m2.fullInfo())
}

func compare1(p book, q book) bool {
	if p.name == q.name && p.author == q.author {
		p.name = "some text" // doest change original variable
		return true
	}
	return false
}

func compare2(p *book, q *book) bool {
	if p.name == q.name && p.author == q.author {
		p.name = "some text" // change original variable
		return true
	}
	return false
}

func example2() {
	var emp employee

	emp.name = "Viktor"
	emp.job = "Go Programmer"
	emp.lastLoggedIn = time.Now().Format(time.RFC850)
	emp.DOB = date.Today()

	fmt.Println(emp)

	p := &emp

	p.name = "Jack"
	emp.job = "Go Expert"

	fmt.Println(*p)
}

func example1() {
	player1 := player{"Messi", "Soccer", 32}
	fmt.Println("player1", player1)
	fmt.Printf("(1) name=%s age=%d\n", player1.name, player1.age)

	player2 := player{
		"Federar",
		"tennis",
		43,
	}

	fmt.Println("player2", player2)
	fmt.Printf("(1) name=%s age=%d\n", player2.name, player2.age)

	var player3 player
	fmt.Println("player 3=", player3)

	player3.name = "Bolt"
	player3.sport = "Sprinter"
	player3.age = 31

	fmt.Println("player 3 updated=", player3)

	player4 := player{
		name:  "Michael Jordan",
		sport: "basketball",
	}

	fmt.Println("player 4=", player4)

	player5 := new(player)
	player5.name = "Tiger Woods"
	player5.sport = "golf"
	player5.age = 44

	fmt.Printf("*player5=%v \n", player5)
	fmt.Printf("*player5=%+v \n", player5)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
