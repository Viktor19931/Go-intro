package animals

// Speaker
type Speaker interface {
	Speaks() string
}

// Perform
func Perform(s Speaker) string{ return s.Speaks() }

// Lion 
type Lion struct{}

// Speaks
func (s Lion) Speaks() string { return "roars" }

// Cat
type Cat struct {}

// Speaks
func (s Cat) Speaks() string { return "meaws" }

// Human
type Human struct {}

// Speaks
func(s Human) Speaks() string { return "talks"}