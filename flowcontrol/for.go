package flowcontrol

import "fmt"

func normalFor() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

func doubleFor() {
	for x := 0; x <= 12; x++ {
		fmt.Println("x: ", x)
		for y := 0; y < 60; y++ {
			fmt.Print(" ", y)
		}
		fmt.Println()
	}
}

func whileFor() {
	x := 0

	for x < 10 {
		fmt.Println("x is minor than 10")
		x++
	}
}

func forWithBreak() {
	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("x is minor than 10")
		} else {
			fmt.Println("x is bigger than 10, i'm out")
			break
		}
	}
}

func forWithContinue() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}

func forAscii() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d - %v\n", i, fmt.Sprintf("%d", i))
	}
}
