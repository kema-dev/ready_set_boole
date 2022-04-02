package main

import (
	"fmt"
	"operation"
)

func main() {
	a := int32(10)
	b := int32(20)
	fmt.Println("adder(a, b) =", operation.Adder(a, b))
}
