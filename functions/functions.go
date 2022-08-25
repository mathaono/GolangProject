package functions

import "fmt"

func Add(n1 int64, n2 int64) int64 {
	return n1 + n2
}

var Funcs = func(txt string) {
	fmt.Println(txt)
}

func Calculations(n1, n2 int64) (int64, int64) {
	sum := n1 + n2
	subtract := n1 - n2
	return sum, subtract
}
