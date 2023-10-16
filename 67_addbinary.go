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
