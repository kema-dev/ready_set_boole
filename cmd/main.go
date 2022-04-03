package main

import (
	operation "RSB/pkg"
	"fmt"
)

func main() {
	a := int32(10)
	b := int32(20)
	fmt.Printf("add(%d, %d) = %d\n", a, b, operation.Add(a, b))
	fmt.Printf("substract(%d, %d) = %d\n", a, b, operation.Substract(a, b))
	fmt.Printf("multiply(%d, %d) = %d\n", a, b, operation.Multiply(a, b))
	fmt.Printf("GrayCode(%d) = %d\n", a, operation.GrayCode(a))
}
