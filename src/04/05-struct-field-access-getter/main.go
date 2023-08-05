package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter/person"
)

func main() {
	p := person.Person{} // boş bir kopya (instance)

	p.FirstName = "Uğur"
	p.LastName = "Özyılmazel"

	fmt.Printf("%+v\n", p) // {FirstName:Uğur LastName:Özyılmazel secret:}

	p.SetSecret("<secret>")

	fmt.Printf("%+v\n", p)  // {FirstName:Uğur LastName:Özyılmazel secret:<secret>}
	fmt.Println(p.Secret()) // <secret>
}
