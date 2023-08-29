package main

import (
	"fmt"
	"unsafe"
)

// Person struct.
type Person struct {
	Name  string
	age   int
	money float64
}

func main() {
	person := Person{Name: "John", age: 30, money: 100.0}

	// Cast name to unsafe pointer
	namePtr := unsafe.Pointer(&person.Name) // nolint
	nameSize := unsafe.Sizeof(person.Name)  // nolint
	fmt.Println("namePtr         :", namePtr)
	fmt.Println("nameSize        :", nameSize)

	// Cast age to unsafe pointer
	agePtr := unsafe.Pointer(&person.age) // nolint
	ageSize := unsafe.Sizeof(person.age)  // nolint
	fmt.Println("agePtr          :", agePtr)
	fmt.Println("ageSize         :", ageSize)

	moneyPtr := unsafe.Pointer(&person.money) // nolint
	moneySize := unsafe.Sizeof(person.money)  // nolint
	fmt.Println("moneyPtr        :", moneyPtr)
	fmt.Println("moneySize       :", moneySize)

	// Get age pointer from name pointer
	foundAge := (*int)(unsafe.Add(namePtr, nameSize))
	fmt.Println("finding agePtr  :", foundAge)
	*foundAge = 10

	// Get money pointer from name pointer
	foundMoney := (*float64)(unsafe.Add(agePtr, ageSize))
	// foundMoney := (*float64)(unsafe.Add(namePtr, nameSize+ageSize)) // also works
	fmt.Println("finding moneyPtr:", foundMoney)
	*foundMoney = 5000.0

	fmt.Println("modified person :", person) // {John 10}
}
