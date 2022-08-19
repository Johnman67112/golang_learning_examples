package datagrouping

import "fmt"

func normalMap() {
	friends := map[string]int{
		"alfred":   5551234,
		"adrienne": 9996674,
	}

	fmt.Println(friends)
	fmt.Println(friends["alfred"])

	friends["gopher"] = 4444444
	fmt.Println(friends)
	fmt.Println(friends["gopher"])
}

func commaOk() {
	friends := map[string]int{
		"alfred":   5551234,
		"adrienne": 9996674,
	}

	thereIs, ok := friends["ghost"]
	fmt.Println(thereIs, ok)

	if thereIs, ok = friends["ghost"]; !ok {
		fmt.Println("there is not")
	} else {
		fmt.Println(thereIs)
	}
}

func rangeMap() {
	anything := map[int]string{
		123: "bla",
		98:  "ble",
		983: "bli",
		18:  "blo",
	}
	fmt.Println(anything)

	for key, value := range anything {
		fmt.Println(key, value)
	}
}

func deleteMapElement() {
	anything := map[int]string{
		123: "bla",
		98:  "ble",
		983: "bli",
		18:  "blo",
	}
	fmt.Println(anything)

	delete(anything, 123)
	fmt.Println(anything)
}
