package athletes

import "strings"

// Info struct
type Info struct {
	Country   string
	HairColor string
}

// Player Struct
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// ToLowercase func
func (p *Player) ToLowercase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.HairColor = strings.ToLower(p.HairColor)
	p.Country = strings.ToLower(p.Country)

	return p
}
