package main

func main() {

	println("For padrão")
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("---------------------------")
	println("For com range")
	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	println("---------------------------")
	println("For tipo while")
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	// for {
	// 	println("Hello, World!")
	// }
}
