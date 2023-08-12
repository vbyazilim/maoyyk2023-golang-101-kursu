package main

import (
	"fmt"
	"reflect"
	"strings"
)

// User holds user model data.
type User struct {
	FirstName string `case:"upper"`
	LastName  string `case:"lower"`
	Age       int    `case:"lower"` // won't affect
}

// Set sets case tag declarations.
// "case" tag is only operational for strings!
func (u *User) Set() {
	v := reflect.Indirect(reflect.ValueOf(u))
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		k := field.Type().Kind()
		if k == reflect.String {
			tagCase, ok := t.Field(i).Tag.Lookup("case")
			if !ok {
				continue
			}

			switch tagCase {
			case "upper":
				if field.String() != "" {
					field.SetString(strings.ToUpper(field.String()))
				}
			case "lower":
				if field.String() != "" {
					field.SetString(strings.ToLower(field.String()))
				}
			}
		}
	}
}

func main() {
	u1 := User{
		FirstName: "Uğur",
		LastName:  "Özyılmazel",
		Age:       51,
	}

	fmt.Printf("%+v\n", u1) // {FirstName:Uğur LastName:Özyılmazel Age:49}

	u1.Set()
	fmt.Println(u1.FirstName) // UĞUR
	fmt.Println(u1.LastName)  // özyılmazel
	fmt.Println(u1.Age)       // 51
}
