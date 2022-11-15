package main

import "fmt"

type Conta struct {
	saldo float64
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}
	fmt.Printf("Simulação: %.2f\n", conta.simular(200))
	fmt.Printf("Saldo: %.2f\n", conta.saldo)
}
