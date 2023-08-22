package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func checkType(d any) {
	switch v := reflect.ValueOf(d); v.Kind() {
	case reflect.Struct:
		s := fmt.Sprintf("%+v", v)
		fmt.Printf("%-24v (%v)\n", s, v.Type())
	case reflect.String:
		fmt.Printf("%-24v (%v)\n", v, v.Type())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %d\n", v, t, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %d\n", v, t, v.Uint())
	case reflect.Bool:
		fmt.Printf("%-24v (%v)\n", v, v.Type())
	case reflect.Float32, reflect.Float64:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v -> %f\n", v, t, v.Float())
	case reflect.Complex64, reflect.Complex128:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Func:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v\n", v, t)
	case reflect.Chan:
		t := "(" + v.Type().String() + ")"
		fmt.Printf("%-24v %-10v\n", v, t)
	case reflect.Map:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Array, reflect.Slice:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Uintptr:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.UnsafePointer:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Pointer:
		t := "(" + v.Type().String() + ")"
		s := fmt.Sprintf("%v", v)
		el := reflect.TypeOf(d).Elem()
		if el.Kind() == reflect.Interface {
			t = "(" + el.Kind().String() + ")"
		}
		fmt.Printf("%-24v %-10v\n", s, t)
	case reflect.Interface:
		fmt.Println("this is not possible")
	case reflect.Invalid:
		fmt.Printf("%-24v\n", v)
	default:
		fmt.Println(v, "unknown")
	}
}

func main() {
	i := 1

	var err error

	data := []any{
		"hello",
		2023,
		int8(127),                    // upper limit
		int16(32767),                 // upper limit
		int32(2147483647),            // upper limit
		int64(9223372036854775807),   // upper limit
		uint(9223372036854775807),    // upper limit
		uint8(255),                   // upper limit
		uint16(65535),                // upper limit
		uint32(4294967295),           // upper limit
		uint64(18446744073709551615), // upper limit
		true,
		false,
		float32(1.0),
		float64(1.0),
		complex(float32(3.0), float32(4.0)),
		complex(5, 7),
		func() {},
		make(chan int),
		map[string]string{"key": "value"},
		[1]string{"array"},
		[]string{"slice"},
		uintptr(unsafe.Pointer(&i)), // nolint
		&i,
		unsafe.Pointer(&i), // nolint
		&err,
		struct{}{},
		nil,
	}

	for _, v := range data {
		checkType(v)
	}
}

// hello                    (string)
// 2023                     (int)      -> 2023
// 127                      (int8)     -> 127
// 32767                    (int16)    -> 32767
// 2147483647               (int32)    -> 2147483647
// 9223372036854775807      (int64)    -> 9223372036854775807
// 9223372036854775807      (uint)     -> 9223372036854775807
// 255                      (uint8)    -> 255
// 65535                    (uint16)   -> 65535
// 4294967295               (uint32)   -> 4294967295
// 18446744073709551615     (uint64)   -> 18446744073709551615
// true                     (bool)
// false                    (bool)
// 1                        (float32)  -> 1.000000
// 1                        (float64)  -> 1.000000
// (3+4i)                   (complex64)
// (5+7i)                   (complex128)
// 0x102b45b50              (func())
// 0x140001000c0            (chan int)
// map[key:value]           (map[string]string)
// [array]                  ([1]string)
// [slice]                  ([]string)
// 1374390648856            (uintptr)
// 0x14000110018            (*int)
// 0x14000110018            (unsafe.Pointer)
// 0x14000102020            (interface)
// {}                       (struct {})
// <invalid reflect.Value>
