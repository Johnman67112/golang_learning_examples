package datagrouping

import "fmt"

func normalArray() {
	//declaration
	var x [5]int
	var y [6]int

	//attribution
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)

	y[0] = 2
	y[1] = 20
	fmt.Println(y[0], y[1])
	fmt.Println(y)

	//type
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//size
	fmt.Println(len(x))
	fmt.Println(len(y))
}

func directDeclarationArray() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
}

func sliceComparingToArray() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//array2 := append(array, 6)
	//fmt.Println(array2)
	slice2 := append(slice, 6)
	fmt.Println(slice2)
}
