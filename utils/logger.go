package utils

import "fmt"

func Logger(message string) {
	divider()
	fmt.Println("---------------- ", message, " ----------------")
	divider()
}

func PlusLogger(message string) {
	divider()
	fmt.Println("++++++ ", message, " ++++++")
	divider()
}

func divider() {
	fmt.Println()
}
