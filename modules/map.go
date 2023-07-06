package modules

import "fmt"

func Map() {
	// map
	var subjects = map[string]float32{"Golang": 23, "Java": 23.3, "Python": 30.0}
	fmt.Println("Map containing subject", subjects)

	// access map value by key
	fmt.Println("Access map value by key", subjects["Golang"])

	// add a new element to a map
	subjects["Javascript"] = 2093.0

	// modified a value in a map
	subjects["Golang"] = 3883.93

	// delete a Java in a map
	delete(subjects, "Java")

	fmt.Println("subjects: ", subjects)

	// loop through map
	for k, v := range subjects {
		fmt.Println(k, ":", v)
	}
}
