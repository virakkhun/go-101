package modules

import "fmt"

func Struct() {
	// struct
	type Person struct {
		name   string
		age    int
		gender string
	}

	var person Person = Person{"Dara", 20, "M"}
	var listOfPersons = [3]Person{{"Dara", 21, "G"}, {"Virak", 20, "M"}, {"Nani", 21, "M"}}

	fmt.Println("Person: ", person)
	fmt.Println("List of person", listOfPersons)

	// function type in struct
	type Rect func(int, int) int

	type rectPara struct {
		length  int
		breadth int
		color   string
		rect    Rect
	}

	result := rectPara{
		length:  10,
		breadth: 20,
		color:   "red",
		rect: func(i1, i2 int) int {
			return i1 * i2
		},
	}

	fmt.Println("Color of rect", result.color)
	fmt.Println("Area of rect", result.rect(result.length, result.breadth))
}
