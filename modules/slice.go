package modules

import "fmt"

func Slice() {
	// slice
	var numSlices = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice of number", numSlices)

	// create a slice from an array
	var sliceFromArray = Peoples[0:2]
	fmt.Println("Slice from an array", sliceFromArray)

	// slice - append
	sliceFromArray = append(sliceFromArray, "Vercel", "AWS")
	fmt.Println("Append the slice", sliceFromArray)

	// slice - copy
	var sliceFromArrayCopied = copy([]string{"Akamai"}, sliceFromArray)
	fmt.Println("Copy the slice", sliceFromArray, sliceFromArrayCopied)
}
