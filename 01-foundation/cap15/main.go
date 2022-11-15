package main

func main() {
	a := 10
	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	println(a)
	b := &a
	println(*b)
}
