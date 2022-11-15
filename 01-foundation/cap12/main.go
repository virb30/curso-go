package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	vinicius := Cliente{
		Nome:  "Vinicius",
		Idade: 30,
		Ativo: true,
	}
	vinicius.Ativo = false
	vinicius.Logradouro = "Bauru"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", vinicius.Nome, vinicius.Idade, vinicius.Ativo)
}
