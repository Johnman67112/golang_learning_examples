package flowcontrol

import "fmt"

func normalIf() {
	x := 2

	if x == 2 {
		fmt.Println("x is equal to 2")
	}
}

func normalElse() {
	x := 3

	if x == 2 {
		fmt.Println("x is equal to 2")
	} else {
		fmt.Println("x is different of 2")
	}
}

func elseIf() {
	x := 3

	if x == 2 {
		fmt.Println("x is equal to 2")
	} else if x == 3 {
		fmt.Println("x is equal to 3")
	}
}

func elseIfElse() {
	x := 3

	if x == 2 {
		fmt.Println("x is equal to 2")
	} else if x == 3 {
		fmt.Println("x is equal to 3")
	} else {
		fmt.Println("x is different of 2 and 3")
	}
}
