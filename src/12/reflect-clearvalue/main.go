package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// Foo1 represents dummy type.
type Foo1 struct {
	A int
	B string
}

// Foo2 represents dummy type.
type Foo2 struct {
	C int
	D string
}

// Foo3 represents dummy type.
type Foo3 struct {
	X bool
}

var (
	errInvalidKind  = errors.New("invalid kind")
	errInvalidValue = errors.New("invalid value")
)

func resetStringValueOfGivenField(d any, f string) error {
	v := reflect.ValueOf(d) // value of d
	k := v.Kind()           // kind of d

	if k != reflect.Ptr {
		return fmt.Errorf("%w, must be a pointer to struct, not a %q", errInvalidKind, k)
	}

	structValue := v.Elem()
	fieldValue := structValue.FieldByName(f)
	if !fieldValue.IsValid() {
		return fmt.Errorf("%w", errInvalidValue)
	}
	if fieldValue.Kind() != reflect.String {
		return fmt.Errorf("%w, %s value must be a string, not a %q", errInvalidKind, f, fieldValue.Kind())
	}
	fieldValue.SetString("")
	return nil
}

func main() {
	foo1 := Foo1{1, "hello"}
	fmt.Printf("foo1: %+v\n", foo1) // foo1: {A:1 B:hello}

	if err := resetStringValueOfGivenField(&foo1, "B"); err != nil {
		log.Print(err)
	}
	fmt.Printf("foo1 after: %+v\n", foo1) // foo1 after: {A:1 B:}

	foo2 := Foo2{2, "world"}
	fmt.Printf("foo2: %+v\n", foo2) // foo2: {C:2 D:world}

	if err := resetStringValueOfGivenField(&foo2, "D"); err != nil {
		log.Print(err)
	}
	fmt.Printf("foo2 after: %+v\n", foo2) // foo2 after: {C:2 D:}

	foo3 := Foo3{}
	if err := resetStringValueOfGivenField(&foo3, "X"); err != nil {
		log.Print(err) // invalid kind, X value must be a string, not a "bool"
	}
	fmt.Printf("foo3 after: %+v\n", foo3)

	// passing non-pointer
	if err := resetStringValueOfGivenField(foo3, "X"); err != nil {
		log.Print(err) // invalid kind, must be a pointer to struct, not a "struct"
	}
}
