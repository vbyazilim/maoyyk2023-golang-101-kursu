package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User holds user data.
type User struct {
	Name      string `json:"name"`
	Email     string
	Age       int
	Optionals map[string]any `json:"-"`
}

func main() {
	incoming := `{
		"name": "Uğur",
		"email": "vigo@example.com",
		"age": 51,
		"foo": 1,
		"bar": "2"
	}`

	u := User{}
	if err := json.Unmarshal([]byte(incoming), &u.Optionals); err != nil {
		log.Fatal(err)
	}

	fmt.Println("u.Optionals", u.Optionals)
	// u.Optionals map[age:51 bar:2 email:vigo@example.com foo:1 name:Uğur]

	if v, ok := u.Optionals["name"].(string); ok {
		u.Name = string(v)
		delete(u.Optionals, "name")
	}
	if v, ok := u.Optionals["email"].(string); ok {
		u.Email = string(v)
		delete(u.Optionals, "email")
	}
	if v, ok := u.Optionals["age"].(float64); ok {
		u.Age = int(v)
		delete(u.Optionals, "age")
	}

	fmt.Printf("u: %+v\n", u)
	// u: {Name:Uğur Email:vigo@example.com Age:51 Optionals:map[bar:2 foo:1]}

	if u.Email != "" {
		fmt.Println("u.Email", u.Email)
		// u.Email vigo@example.com
	}
	if u.Age != 0 {
		fmt.Println("u.Age", u.Age)
		// u.Age 51
	}

	if len(u.Optionals) > 0 {
		fmt.Printf("you have %d invalid field(s)\n", len(u.Optionals))
		// you have 2 invalid field(s)

		for v := range u.Optionals {
			fmt.Printf("%q\n", v)
		}
		// "foo"
		// "bar"
	}
}
