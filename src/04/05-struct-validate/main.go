package main

import (
	"fmt"
	"log"

	"github.com/isacikgoz/defaults"
)

// User holds basic user information.
type User struct {
	Name     string `validate:"notempty"`
	Email    string `validate:"email"`
	Homepage string `validate:"url"      default:"https://vbyazilim.com"`
}

func main() {
	var u User

	// set defaults
	if err := defaults.Set(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u) // {Name: Email: Homepage:https://vbyazilim.com}

	u.Name = "Uğur Özyılmazel"
	u.Email = "ugur@fake.com"
	// u.Email = ""
	// u.Homepage = ""
	// u.Homepage = "foooo"

	if err := defaults.Validate(&u); err != nil {
		log.Fatal(err)
	}
}
