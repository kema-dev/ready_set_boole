package operation

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
