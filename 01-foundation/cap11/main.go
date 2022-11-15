package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	vinicius := Cliente{
		Nome:  "Vinicius",
		Idade: 30,
		Ativo: true,
	}
	vinicius.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", vinicius.Nome, vinicius.Idade, vinicius.Ativo)
}
