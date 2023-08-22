package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// GradeType defines generic grade type.
type GradeType interface {
	constraints.Ordered
}

// AgeType defines generic age type.
type AgeType interface {
	constraints.Ordered
}

// Student represents generic student type model.
type Student[gradeType GradeType, ageType AgeType] struct {
	Name  string
	Age   gradeType
	Grade ageType
}

func main() {
	student := Student[int, float64]{
		Name:  "John",
		Age:   20,
		Grade: 10.21,
	}

	fmt.Printf("%+v\n", student) // {Name:John Age:20 Grade:10.21}
}
