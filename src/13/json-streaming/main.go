package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Person represents person model.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// People represents collection on Person slice.
type People []Person

func main() {
	// bulk data, looks like json array?
	incoming := `
	{"name": "Fred", "age": 40}
	{"name": "Mary", "age": 21}
	{"name": "Pat", "age": 30}
	`

	decoder := json.NewDecoder(strings.NewReader(incoming))

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)

	var p Person
	var people People

	for decoder.More() {
		if err := decoder.Decode(&p); err != nil {
			log.Print("decode err", err)
			continue
		}

		fmt.Printf("p: %+v\n", p)
		// p: {Name:Fred Age:40}
		// p: {Name:Mary Age:21}
		// p: {Name:Pat Age:30}

		people = append(people, p)

		if err := encoder.Encode(p); err != nil {
			log.Panic("encode err", err) // do not panic!
		}
	}

	fmt.Println(b.Bytes())
	// [123 34 110 97 109 101 34 58 34 70 114 101 100 34 44 34 97 103 101 34 58 52 48 125 10 123 34 110 97 109 101 34 58 34 77 97 114 121 34 44 34 97 103 101 34 58 50 49 125 10 123 34 110 97 109 101 34 58 34 80 97 116 34 44 34 97 103 101 34 58 51 48 125 10]

	// fmt.Println(string(b.Bytes()))
	fmt.Println(b.String())
	// {"name":"Fred","age":40}
	// {"name":"Mary","age":21}
	// {"name":"Pat","age":30}

	fmt.Printf("people: %+v\n", people)
	// people: [{Name:Fred Age:40} {Name:Mary Age:21} {Name:Pat Age:30}]

	j, _ := json.MarshalIndent(people, "", "    ")
	fmt.Printf("%s\n", j)
	// [
	//     {
	//         "name": "Fred",
	//         "age": 40
	//     },
	//     {
	//         "name": "Mary",
	//         "age": 21
	//     },
	//     {
	//         "name": "Pat",
	//         "age": 30
	//     }
	// ]
}
