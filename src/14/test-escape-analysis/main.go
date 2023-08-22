package main

import "fmt"

// User holds user data.
type User struct {
	Email    string
	FullName string
}

// NewUserAsValue creates new User instance as value semantics.
func NewUserAsValue(email, fullName string) User {
	return User{email, fullName}
}

// NewUserAsPointer creates new User instance as pointer semantics.
func NewUserAsPointer(email, fullName string) *User {
	return &User{email, fullName}
}

func main() {
	u1 := NewUserAsValue("vigo@me.com", "Uğur Özy")
	u2 := NewUserAsPointer("vigo@me.com", "Uğur Özy")

	fmt.Println("u1 -> value", u1)   // u1 -> value {vigo@me.com Uğur Özy}
	fmt.Println("u2 -> pointer", u2) // u2 -> pointer &{vigo@me.com Uğur Özy}
}
