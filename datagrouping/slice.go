package datagrouping

import "fmt"

func normalSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
}

func changeSliceItemValue() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
}

func forRangeSlice() {
	slice := []string{"banana", "apple", "strawberry"}

	for i, v := range slice {
		fmt.Println("i =", i, "| v =", v)
	}
}

func appendSlice() {
	slice := []string{"banana", "apple", "strawberry"}

	slice = append(slice, "watermelon")
	for i, v := range slice {
		fmt.Println("i =", i, "| v =", v)
	}
}

func sliceSlice() {
	flavors := []string{"pepperoni", "mozzarela", "pineapple", "marguerita"}

	slice := flavors[0:2]
	fmt.Println(slice)
}

func sliceAllItems() {
	flavors := []string{"pepperoni", "mozzarela", "pineapple", "marguerita"}

	slice := flavors[:]
	slice2 := flavors[0:]
	slice3 := flavors[:4]
	slice4 := flavors[0:4]

	fmt.Printf("%v\n%v\n%v\n%v\n%v", slice, slice2, slice3, slice3, slice4)
}

func forAllSliceItems() {
	flavors := []string{"pepperoni", "mozzarela", "pineapple", "marguerita"}
	for i := 0; i < len(flavors); i++ {
		fmt.Println(flavors[i])
	}
}

func removeSliceItems() {
	flavors := []string{"pepperoni", "mozzarela", "pineapple", "marguerita"}

	flavors = append(flavors[:2], flavors[3:]...)

	fmt.Println(flavors)
}

func appendUsage() {
	aSlice := []int{1, 2, 3, 4}
	anotherSlice := []int{9, 10, 11, 12}

	aSlice = append(aSlice, 5, 6, 7, 8)
	fmt.Println(aSlice)

	aSlice = append(aSlice, anotherSlice...)
	fmt.Println(aSlice)
}

func sliceMake() {
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 11)
	fmt.Println(slice, len(slice), cap(slice))
}

func multiDimensionalSlice() {
	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(ss)

	sss := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18},
		},
		{
			{19, 20, 21},
			{22, 23, 24},
			{25, 26, 27},
		},
	}

	fmt.Println(sss)
}

func underlyingArray() {
	firstSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(firstSlice)

	secondSlice := append(firstSlice[0:2], firstSlice[4:]...)

	fmt.Println(secondSlice)
	fmt.Println(firstSlice)
}
