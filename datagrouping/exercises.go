package datagrouping

import "fmt"

func ex1() {
	array := [5]int{1, 2, 3, 4, 5}

	for i, v := range array {
		fmt.Println("i:", i, "v:", v)
	}

	fmt.Printf("%T", array)
}

func ex2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println("i:", i, "v:", v)
	}

	fmt.Printf("%T", slice)
}

func ex3() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}

func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func ex6() {
	slice := make([]string, 26)
	slice[0], slice[1], slice[2], slice[3], slice[4],
		slice[5], slice[6], slice[7], slice[8], slice[9],
		slice[10], slice[11], slice[12], slice[13], slice[14],
		slice[15], slice[16], slice[17], slice[18], slice[19],
		slice[20], slice[21], slice[22], slice[23], slice[24],
		slice[25] = "Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
		"Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo",
		"Sergipe", "Tocantins"

	//brazilian states lol

	fmt.Println(slice[:])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func ex7() {
	slice := [][]string{
		{"Jonas", "Almeida", "Soccer"},
		{"Melis", "Jackson", "Tenis"},
		{"Carlo", "Johnson", "Basketball"},
	}

	for _, v := range slice {
		fmt.Println("First Name:", v[0], "| Last Name:", v[1], "| Hobby:", v[2])
	}
}

func ex8() {
	mappy := map[string]string{
		"Jonas_Almeida": "Soccer",
		"Melis_Jackson": "Tenis",
		"Carlo_Johnson": "Basketball",
	}

	for k, v := range mappy {
		fmt.Println(k, v)
	}

}

func ex9() {
	mappy := map[string]string{
		"Jonas_Almeida": "Soccer",
		"Melis_Jackson": "Tenis",
		"Carlo_Johnson": "Basketball",
	}

	mappy["Vitor_Doommed"] = "World_Domination"

	for k, v := range mappy {
		fmt.Println(k, v)
	}
}

func ex10() {
	mappy := map[string]string{
		"Jonas_Almeida": "Soccer",
		"Melis_Jackson": "Tenis",
		"Carlo_Johnson": "Basketball",
		"Vitor_Doommed": "World_Domination",
	}

	delete(mappy, "Vitor_Doommed")

	for k, v := range mappy {
		fmt.Println(k, v)
	}
}
