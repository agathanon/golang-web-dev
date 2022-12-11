package main

import "fmt"

// person represents a person
type person struct {
	// firstName is a person's first name
	firstName string

	// lastName is a person's last name
	lastName string
}

// pSpeak prints a phrase a person says
func (p person) pSpeak() {
	fmt.Println("Yes, hello! My name is", p.firstName, p.lastName)
}

// secretAgent represents a person who is a secret agent
type secretAgent struct {
	// undercover is a boolean flag that represents an agent's undercover
	// status.
	undercover bool

	// Embedded person type
	person
}

// saSpeak prints a phrase a secretAgent says
func (sa secretAgent) saSpeak() {
	fmt.Println("I'm definitely not a cop.")
}

func main() {
	p := person{
		firstName: "John",
		lastName: "Doe",
	}
	fmt.Printf("p.firstName: %s\n", p.firstName)
	p.pSpeak()

	s := secretAgent{
		undercover: true,
		person:     person{
			firstName: "Ray",
			lastName: "Dacted",
		},
	}
	fmt.Printf("s.undercover? %v\n", s.undercover)
	s.saSpeak()
	s.pSpeak()
}
