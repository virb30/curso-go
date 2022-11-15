package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Println("A empresa foi desativada")
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	vinicius := Cliente{
		Nome:  "Vinicius",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativacao(vinicius)
	Desativacao(minhaEmpresa)
}
