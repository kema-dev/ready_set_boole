package operation

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(a int32, b int32) int32 {
	sum := a ^ b
	carry := (a & b) << 1
	for i := 0; i < 32; i++ {
		tmp := (sum & carry) << 1
		sum = sum ^ carry
		carry = tmp
	}
	return sum | carry
}

func Substract(a int32, b int32) int32 {
	return Add(a, -b)
}

func Multiply(a int32, b int32) int32 {
	var res int32 = 0
	for (b != 0) {
		b = Add(b, -1)
		res = Add(res, a)
	}
	return res
}

func GrayCode(a int32) int32 {
	return a ^ (a >> 1)
}

func BoolEval(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]bool, 0)
	for _, c := range s {
		switch c {
		case '0':
			stack = append(stack, false)
		case '1':
			stack = append(stack, true)
		case '!':
		if len(stack) < 1 {
			return false
		}
		b1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, !b1)
		case '&':
			if len(stack) < 2 {
				return false
			}
			b1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b1 && b2)
		case '|':
			if len(stack) < 2 {
				return false
			}
			b1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b1 || b2)
		case '^':
			if len(stack) < 2 {
				return false
			}
			b1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b1 != b2)
		case '>':
			if len(stack) < 2 {
				return false
			}
			b1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, !(b1 && !b2))
		case '=':
			if len(stack) < 2 {
				return false
			}
			b1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, (b1 && b2) || (!b1 && !b2))
		default:
			fmt.Println("invalid character: ", string(c))
			return false
		}
	}
	if len(stack) != 1 {
		fmt.Println("stack length is not 1 (", len(stack), ") stack:", stack)
		return false
	}
	return stack[0]
}

func TruthTable(s string) {
	if len(s) == 0 {
		return
	}
	str := ""
	for _, c := range s {
		if (!(c >= 'A' && c <= 'Z') && (c != '!') && (c != '&') && (c != '|') && (c != '^') && (c != '>') && (c != '=')) {
			fmt.Println("Invalid character: ", string(c))
			return
		} else {
			str += string(c)
		}
	}
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charset := ""
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(letters); j++ {
			if str[i] == letters[j] {
				charset += string(str[i])
				break
			}
		}
	}
	for i := 0; i < len(charset); i++ {
		fmt.Print("| " + string(charset[i]) + " ")
	}
	fmt.Println("|  =  |")
	for i := 0; i < len(charset); i++ {
		fmt.Print("|---")
	}
	fmt.Println("|-----|")
	max := 1 << uint(len(charset))
	for i := 0; i < max; i++ {
		conv := strconv.FormatInt(int64(i), 2)
		for len(conv) < len(charset) {
			conv = conv + "0"
		}
		buf := str
		for j := 0; j < len(charset); j++ {
			buf = strings.ReplaceAll(buf, string(charset[j]), string(conv[j]))
		}
		for i := 0; i < len(charset); i++ {
			fmt.Print("| " + string(buf[i]) + " ")
		}
		fmt.Printf("|%-5t|\n", BoolEval(buf))
	}
}
