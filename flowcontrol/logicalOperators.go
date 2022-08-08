package flowcontrol

import (
	"fmt"
)

func orOperator() {
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("x is 2 or 3")
	}
}

func andOperator() {
	x := 2
	y := 3

	if x == 2 && y == 3 {
		fmt.Println("x is 2 and y is 3")
	}
}

func notOperator() {
	x := 2
	y := 3

	if !(x == 2 || x == 3) {
		fmt.Println("x is 2 or 3")
	}

	if !(x == 2 && y == 3) {
		fmt.Println("x is 2 and y is 3")
	}
}
