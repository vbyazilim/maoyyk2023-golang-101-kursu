package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User holds user model data. User struct has json tags!
type User struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	BirthDate time.Time  `json:"birth_date"`
	Admin     bool       `json:"admin"`
	LastVisit *time.Time `json:"last_visit"`
}

// Users holds user slice of Users.
type Users []User

func main() {
	input := `[
		{
			"first_name": "Uğur",
			"last_name": "Özyılmazel",
			"age": 51,
			"birth_date": "1972-08-13T09:00:00.0+03:00"
		},
		{
			"first_name": "Ömer",
			"last_name": "Özyılmazel",
			"age": 41,
			"birth_date": "1982-08-13T14:00:00.0+03:00",
			"admin": true
		},
		{
			"this": "This",
			"is": "is",
			"fake": "fake"
		}
	]`

	b := []byte(input)
	d := json.NewDecoder(bytes.NewReader(b))
	// accepts io.Reader, convert byte slice -> io.Reader satisfier

	d.DisallowUnknownFields()
	// raise error for unknown fields

	var users Users

	// if err := json.Unmarshal(b, &users); err != nil {
	// 	log.Fatal(err)
	// }

	if err := d.Decode(&users); err != nil {
		fmt.Println(err) // json: unknown field "this"
	}

	fmt.Printf("%+v\n", users)
	// [{FirstName:Uğur LastName:Özyılmazel Age:51 BirthDate:1972-08-13 09:00:00 +0300 +0300 Admin:false LastVisit:<nil>} {FirstName:Ömer LastName:Özyılmazel Age:41 BirthDate:1982-08-13 14:00:00 +0300 +03 Admin:true LastVisit:<nil>} {FirstName: LastName: Age:0 BirthDate:0001-01-01 00:00:00 +0000 UTC Admin:false LastVisit:<nil>}]

	for _, user := range users {
		fmt.Printf("%+v\n", user)
		fmt.Println(user.LastVisit, user.LastVisit == nil)
	}
	// {FirstName:Uğur LastName:Özyılmazel Age:51 BirthDate:1972-08-13 09:00:00 +0300 +0300 Admin:false LastVisit:<nil>}
	// <nil> true
	// {FirstName:Ömer LastName:Özyılmazel Age:41 BirthDate:1982-08-13 14:00:00 +0300 +03 Admin:true LastVisit:<nil>}
	// <nil> true
	// {FirstName: LastName: Age:0 BirthDate:0001-01-01 00:00:00 +0000 UTC Admin:false LastVisit:<nil>}
	// <nil> true

	j, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", j)
	// [
	//     {
	//         "first_name": "Uğur",
	//         "last_name": "Özyılmazel",
	//         "age": 51,
	//         "birth_date": "1972-08-13T09:00:00+03:00",
	//         "admin": false,
	//         "last_visit": null
	//     },
	//     {
	//         "first_name": "Ömer",
	//         "last_name": "Özyılmazel",
	//         "age": 41,
	//         "birth_date": "1982-08-13T14:00:00+03:00",
	//         "admin": true,
	//         "last_visit": null
	//     },
	//     {
	//         "first_name": "",
	//         "last_name": "",
	//         "age": 0,
	//         "birth_date": "0001-01-01T00:00:00Z",
	//         "admin": false,
	//         "last_visit": null
	//     }
	// ]
}
