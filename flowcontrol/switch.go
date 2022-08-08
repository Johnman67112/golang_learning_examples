package flowcontrol

import (
	"fmt"
)

func baseSwitch() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("x is minor than 5")
	case x == 5:
		fmt.Println("x is equal to 5")
	case x >= 5:
		fmt.Println("x is bigger than 5")
	}
}

func statementFallthroughSwitch() {
	officeToday := "Marcos"

	switch officeToday {
	case "George":
		fmt.Println("George is on the office today")
	case "Marcos":
		fmt.Println("Marcos is on the office today")
		fallthrough
	case "Gabe":
		fmt.Println("Gabe is on the office today")
	}
}

func defaultSwitch() {
	officeToday := "bla"

	switch officeToday {
	case "George":
		fmt.Println("George is on the office today")
	case "Marcos":
		fmt.Println("Marcos is on the office today")
	case "Gabe":
		fmt.Println("Gabe is on the office today")
	default:
		fmt.Println("Nobody on the office, #failure")
	}
}

func doubleValueSwitch() {
	x := 2

	switch x {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case 5, 6:
		fmt.Println("5 or 6")
	default:
		fmt.Println("Other number idk")
	}
}

func typeSwitch() {
	var x interface{}

	switch x.(type) {
	case int:
		fmt.Println("it's an int")
	case bool:
		fmt.Println("it's a bool")
	case string:
		fmt.Println("it's a string")
	case float64:
		fmt.Println("it's a float")
	default:
		fmt.Printf("%+v\n", x)
	}
}

func initializeSwitch() {
	switch x := 2; {
	case x == 1:
		fmt.Println("1")
	case x == 2:
		fmt.Println("2")
	case x == 3:
		fmt.Println("3")
	case x == 4:
		fmt.Println("4")
	}
}
