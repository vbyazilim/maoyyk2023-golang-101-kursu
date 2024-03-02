package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

// User holds user model data. User struct has json tags!
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	BirthDate time.Time `json:"birth_date"`
	Admin     bool      `json:"admin"`
	LastVisit time.Time `json:"-"` // omitted, ignore
}

const customTimeLayout = "2006-01-02T15:04:05-07:00"

// CustomTime is a custom type definition uses time.Time, uses custom marshal format.
type CustomTime struct {
	time.Time
}

// MarshalJSON marshals CustomTime with using custom time layout.
func (ct CustomTime) MarshalJSON() ([]byte, error) { // nolint
	return []byte(`"` + ct.Time.Format(customTimeLayout) + `"`), nil
}

// UserWithCustomTime holds user model data with custom time type!
type UserWithCustomTime struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	BirthDate time.Time  `json:"birth_date"`
	Admin     bool       `json:"admin"`
	LastVisit CustomTime `json:",omitempty"`
}

// OtherUser holds user model data, only one field has tag!
type OtherUser struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
	Admin     bool
	LastVisit *time.Time `json:",omitempty"` // this field uses pointer, can be nil, can be omitted if unset!
}

func describeStruct(v any) {
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf(
			"Name: %-12s Type: %-18s Tag: %-20s json: %s\n",
			field.Name,
			field.Type,
			field.Tag,
			field.Tag.Get("json"),
		)
	}
	fmt.Println()
}

func marshalStruct(v any) error {
	j, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n\n", j)
	return nil
}

func main() {
	trTZ := time.FixedZone("UTC+3", +3*60*60)
	now := time.Now()

	u1 := User{
		FirstName: "Uğur",
		LastName:  "Özy",
		Age:       49,
		BirthDate: time.Date(1972, time.August, 13, 10, 0, 0, 0, trTZ),
		Admin:     true,
		LastVisit: now,
	}
	describeStruct(u1)
	// Name: FirstName    Type: string             Tag: json:"first_name"    json: first_name
	// Name: LastName     Type: string             Tag: json:"last_name"     json: last_name
	// Name: Age          Type: int                Tag: json:"age"           json: age
	// Name: BirthDate    Type: time.Time          Tag: json:"birth_date"    json: birth_date
	// Name: Admin        Type: bool               Tag: json:"admin"         json: admin
	// Name: LastVisit    Type: time.Time          Tag: json:"-"             json: -

	u2 := OtherUser{
		FirstName: "Ezel",
		LastName:  "Özy",
		Age:       10,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		// we don't have LastVisit field!
	}
	describeStruct(u2)
	// Name: FirstName    Type: string             Tag:                      json:
	// Name: LastName     Type: string             Tag:                      json:
	// Name: Age          Type: int                Tag:                      json:
	// Name: BirthDate    Type: time.Time          Tag:                      json:
	// Name: Admin        Type: bool               Tag:                      json:
	// Name: LastVisit    Type: *time.Time         Tag: json:",omitempty"    json: ,omitempty

	u3 := OtherUser{
		FirstName: "Ezel",
		LastName:  "Özy",
		Age:       12,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		LastVisit: &now,
	}
	describeStruct(u3)
	// Name: FirstName    Type: string             Tag:                      json:
	// Name: LastName     Type: string             Tag:                      json:
	// Name: Age          Type: int                Tag:                      json:
	// Name: BirthDate    Type: time.Time          Tag:                      json:
	// Name: Admin        Type: bool               Tag:                      json:
	// Name: LastVisit    Type: *time.Time         Tag: json:",omitempty"    json: ,omitempty

	u4 := UserWithCustomTime{
		FirstName: "Ezel",
		LastName:  "Özyılmazel",
		Age:       12,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		LastVisit: CustomTime{now},
	}
	describeStruct(u4)
	// Name: FirstName    Type: string             Tag: json:"first_name"    json: first_name
	// Name: LastName     Type: string             Tag: json:"last_name"     json: last_name
	// Name: Age          Type: int                Tag: json:"age"           json: age
	// Name: BirthDate    Type: time.Time          Tag: json:"birth_date"    json: birth_date
	// Name: Admin        Type: bool               Tag: json:"admin"         json: admin
	// Name: LastVisit    Type: main.CustomTime    Tag: json:",omitempty"    json: ,omitempty

	if err := marshalStruct(u1); err != nil {
		log.Fatal(err)
	}
	// {
	//     "first_name": "Uğur",
	//     "last_name": "Özy",
	//     "age": 49,
	//     "birth_date": "1972-08-13T10:00:00+03:00",
	//     "admin": true
	// }

	if err := marshalStruct(u2); err != nil {
		log.Fatal(err)
	}
	// {
	//     "FirstName": "Ezel",
	//     "LastName": "Özy",
	//     "Age": 10,
	//     "BirthDate": "2011-08-13T07:00:00+03:00",
	//     "Admin": false
	// }

	if err := marshalStruct(u3); err != nil {
		log.Fatal(err)
	}
	// {
	//     "FirstName": "Ezel",
	//     "LastName": "Özy",
	//     "Age": 12,
	//     "BirthDate": "2011-08-13T07:00:00+03:00",
	//     "Admin": false,
	//     "LastVisit": "2023-08-13T14:07:09.537756+03:00"
	// }

	if err := marshalStruct(u4); err != nil {
		log.Fatal(err)
	}
	// {
	//     "first_name": "Ezel",
	//     "last_name": "Özyılmazel",
	//     "age": 12,
	//     "birth_date": "2011-08-13T07:00:00+03:00",
	//     "admin": false,
	//     "LastVisit": "2023-08-13T14:07:09+03:00"
	// }
}
