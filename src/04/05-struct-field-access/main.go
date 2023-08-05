package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access/person"
)

func main() {
	p := person.Person{} // boş bir kopya (instance)

	p.FirstName = "Uğur"
	p.LastName = "Özyılmazel"

	fmt.Printf("p: %#v\n", p) // p: person.Person{FirstName:"Uğur", LastName:"Özyılmazel", secret:""}

	// fmt.Println(p.secret) // p.secret undefined (type person.Person has no field or method secret)
}
