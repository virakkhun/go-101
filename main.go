package main

import (
	"go-101/modules"
	"go-101/utils"
)

func main() {

	utils.PlusLogger("Array")
	modules.Array()

	utils.PlusLogger("Slice")
	modules.Slice()

	utils.PlusLogger("Map")
	modules.Map()

	utils.PlusLogger("Struct")
	modules.Struct()

	utils.PlusLogger("Pointer")
	modules.Pointer()
}
