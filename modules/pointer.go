package modules

import (
	"fmt"
	"go-101/utils"
)

func Pointer() {
	var num int = 5

	// print value
	fmt.Println("Value of num ", num)

	// address of num
	fmt.Println("Address of num", &num)

	// create a pointer variable
	var name = "John"
	var namePtr *string

	namePtr = &name

	fmt.Println("Value of namePtr: ", *namePtr)
	fmt.Println("Address of the variable: ", &name)

	utils.PlusLogger("Using Pointer")
	usingPointer()
}

func usingPointer() {
	var num int     // declare a variable
	var numPtr *int // declare a pointer variable

	num = 22
	fmt.Println("Address of num", &num)
	fmt.Println("Value of num", num)

	numPtr = &num
	fmt.Println("Address of numPtr", numPtr)
	fmt.Println("Value of numPtr", *numPtr)

	num = 11
	fmt.Println("Address of numPtr", numPtr)
	fmt.Println("Value of numPtr", *numPtr)

	*numPtr = 2
	fmt.Println("Address of num", &num)
	fmt.Println("Value of num", num)

	var number = 55

	funcPtr(&number)

	fmt.Println("Value of number", number)

	result := returnPtr()

	fmt.Println("Value from func return ptr:", *result)

	ptrWithStruct()
}

func funcPtr(num *int) {
	fmt.Println("arg num:", *num)
	// dereference
	*num = 20
}

func returnPtr() *string {
	message := "Hello Pointer"

	return &message
}

func ptrWithStruct() {
	type Fruit struct {
		weight int
		name   string
	}

	var apple = Fruit{name: "Nez Apple", weight: 20}
	var applePtr *Fruit

	applePtr = &apple

	// struct instance
	fmt.Println(apple)

	// struct type pointer
	fmt.Println(applePtr)

	// access value from applePtr
	fmt.Println("Name:", applePtr.name)
	fmt.Println("Weight:", applePtr.weight)
}
