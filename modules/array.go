package modules

import "fmt"

var Peoples = [...]string{"Apple", "Microsoft", "Amazon", "Google"}

func Array() {
	// fixed size
	var arrayOfInteger = [5]int{1, 2, 3, 4, 5}
	// non-fixed size

	fmt.Println("Array of integer", arrayOfInteger)
	fmt.Println("Size of array", len(arrayOfInteger))

	fmt.Println("Iterate the array")

	// simple loop
	for i := 0; i < len(arrayOfInteger); i++ {
		fmt.Println(arrayOfInteger[i])
	}

	// using with range as an iterator
	for index, item := range Peoples {
		fmt.Println(index, ":", item)
	}
}
