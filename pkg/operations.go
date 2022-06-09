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
	for i := 0; i < 32; i++ {
		if (b & 1) != 0 {
			res = Add(res, a)
		}
		a = a << 1
		b = b >> 1
	}
	return res
}

func GrayCode(a int32) int32 {
	return a ^ (a >> 1)
}

func BoolEval(s string) bool {
	if len(s) == 0 {
		fmt.Println("Provided string is empty")
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
	for _, c := range s {
		if (!(c >= 'A' && c <= 'Z') && (c != '!') && (c != '&') && (c != '|') && (c != '^') && (c != '>') && (c != '=')) {
			fmt.Println("Invalid character: ", string(c))
			return
		}
	}
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charset := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(letters); j++ {
			if s[i] == letters[j] {
				charset += string(s[i])
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
		buf := s
		for j := 0; j < len(charset); j++ {
			buf = strings.ReplaceAll(buf, string(charset[j]), string(conv[j]))
		}
		for i := 0; i < len(charset); i++ {
			fmt.Print("| " + string(buf[i]) + " ")
		}
		fmt.Printf("|%-5t|\n", BoolEval(buf))
	}
}

func NegationNormalForm(s string) string {
	if len(s) == 0 {
			fmt.Println("Provided string is empty")
			return ""
	}
	for _, c := range s {
		if (!(c >= 'A' && c <= 'Z') && (c != '!') && (c != '&') && (c != '|') && (c != '^') && (c != '>') && (c != '=')) {
			fmt.Println("Invalid character: ", string(c))
			return ""
		}
	}
}
