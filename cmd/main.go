package main

import (
	operation "RSB/pkg"
	"fmt"
)

func main() {
	// a := int32(10)
	// b := int32(20)
	// fmt.Printf("add(%d, %d) = %d\n", a, b, operation.Add(a, b))
	// fmt.Printf("substract(%d, %d) = %d\n", a, b, operation.Substract(a, b))
	// fmt.Printf("multiply(%d, %d) = %d\n", a, b, operation.Multiply(a, b))
	// fmt.Printf("graycode(%d) = %d\n", a, operation.GrayCode(a))
	// s := "10&"
	// fmt.Printf("booleval(%s) = %t\n", s, operation.BoolEval(s))
	// s = "10|"
	// fmt.Printf("booleval(%s) = %t\n", s, operation.BoolEval(s))
	// s = "11>"
	// fmt.Printf("booleval(%s) = %t\n", s, operation.BoolEval(s))
	// s = "10="
	// fmt.Printf("booleval(%s) = %t\n", s, operation.BoolEval(s))
	// s = "1011||="
	// fmt.Printf("booleval(%s) = %t\n", s, operation.BoolEval(s))
	// s = "ABCD||="
	// fmt.Printf("truthtable(%s) =\n", s)
	// operation.TruthTable(s)
	s := "AB^CD"
	fmt.Printf("NegationNormalForm(%s) = %s\n", s, operation.NegationNormalForm(s))
}
