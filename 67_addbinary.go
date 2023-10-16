package main

import (
	"fmt"
)

func main() {
	fmt.Println("1:", addBinary("11", "1"))
	fmt.Println("2:", addBinary("1010", "1011"))
	fmt.Println("3:", addBinary("1", "111"))
	fmt.Println("4:", addBinary("1111", "1111"))

}
func addBinary(a string, b string) string {
	l := a
	s := b
	if len(a) < len(b) {
		l = b
		s = a
	}
	res := make([]uint8, 0)
	d := '0'
	koef := len(l) - len(s)
	for i := len(l) - 1; i != -1; i-- {
		x := l[i]
		var y uint8 = '0'
		if i-koef >= 0 {
			y = s[i-koef]
		}
		if x == '1' && y == '1' {
			if d == '1' {
				res = append([]uint8{1}, res...)
			} else {
				res = append([]uint8{0}, res...)
			}
			d = '1'
		}
		if (x == '1' && y == '0') || (x == '0' && y == '1') {
			if d == '1' {
				res = append([]uint8{0}, res...)
			} else {
				d = '0'
				res = append([]uint8{1}, res...)
			}
		}
		if x == '0' && y == '0' {
			if d == '1' {
				res = append([]uint8{1}, res...)
				d = '0'
			} else {
				res = append([]uint8{0}, res...)
			}
		}
	}
	if d == '1' {
		res = append([]uint8{1}, res...)
	}
	var str string
	for _, e := range res {
		if e == 1 {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func addBinaryTheBest(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ans := make([]byte, len(a)+1)
	var carry byte
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i-1, j-1
		var a2Digit, b2Digit byte
		if i >= 0 {
			a2Digit = a[i] - '0'
		}
		if j >= 0 {
			b2Digit = b[j] - '0'
		}
		// sum and carry of full adder
		sum := a2Digit ^ b2Digit ^ carry
		carry = a2Digit&b2Digit | carry&(a2Digit^b2Digit)
		ans[i+1] = sum + '0'
	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
