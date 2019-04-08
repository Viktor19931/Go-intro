package athletes

import "strings"

// Info
type Info struct {
	Country   string
	HairColor string
}

// Player
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// ToLowercase
func (p *Player) ToLowercase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColor = strings.ToLower(p.HairColor)

	return p
}
