package main

func soma(a, b int) int {
	println("sum")
	return a + b
}

func increment(value *int) {
	println("incrementing...")
	*value += 1
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	println(minhaVar1)
	increment(&minhaVar1)
	println(minhaVar1)
	println(soma(minhaVar1, minhaVar2))
	println(minhaVar1)
	increment(&minhaVar1)
	println(minhaVar1)
}
