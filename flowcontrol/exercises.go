package flowcontrol

import "fmt"

func ex1() {
	x := 1
	for {
		fmt.Println(x)
		if x >= 10000 {
			break
		}
		x++
	}
}

func ex2() {
	for x := 65; x <= 90; x++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("%d - %#U\n", x, x)
		}
	}
}

func ex3() {
	x := 2002
	for x <= 2022 {
		fmt.Println(x)
		x++
	}
}

func ex4() {
	x := 2002
	for {
		fmt.Println(x)
		if x >= 2022 {
			break
		}
		x++
	}
}

func ex5() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}
}
func ex6() {
	x := 2

	if x == 2 {
		fmt.Println("x is 2")
	}
}
func ex7() {
	x := 3

	if x == 2 {
		fmt.Println("x is 2")
	} else if x == 3 {
		fmt.Println("x is 3")
	} else {
		fmt.Println("x is not 2 or 3")
	}
}
func ex8() {
	x := 1

	switch {
	case x == 1:
		fmt.Println("x is equal to 1")
	case x == 2:
		fmt.Println("x is equal to 2")
	case x == 3:
		fmt.Println("x is equal to 3")
	default:
		fmt.Println("x is not equal to 1, 2 or 3")
	}
}
func ex9() {
	favoriteSport := "games"

	switch favoriteSport {
	case "footbol":
		fmt.Println("nice kick")
	case "games":
		fmt.Println("nice play")
	case "chambara":
		fmt.Println("nice hit")
	default:
		fmt.Println("nice nothing")
	}
}

func ex10() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
