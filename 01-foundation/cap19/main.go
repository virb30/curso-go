package main

import "fmt"

func main() {
	var minhaVar interface{} = "Viniboscoa"
	println(minhaVar.(string))
	var minhaVar2 interface{} = "15"
	res, ok := minhaVar2.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
