package main

import "fmt"

func main() {

	salarios := map[string]int{"Vinicius": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Vinicius"])
	delete(salarios, "Vinicius")
	fmt.Printf("%v\n", salarios)
	salarios["Vini"] = 5000
	fmt.Println(salarios["Vini"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Vinicius"] = 1000
}
