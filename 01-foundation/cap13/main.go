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

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	vinicius := Cliente{
		Nome:  "Vinicius",
		Idade: 30,
		Ativo: true,
	}
	vinicius.Desativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", vinicius.Nome, vinicius.Idade, vinicius.Ativo)

	vinicius2 := Cliente{
		Nome:  "Vinicius 2",
		Idade: 30,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", vinicius2.Nome, vinicius2.Idade, vinicius2.Ativo)
}
