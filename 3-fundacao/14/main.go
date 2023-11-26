package main

import "fmt"

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) adicionar(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := NewConta()
	//conta := Conta{saldo: 100}
	conta.adicionar(200)
	fmt.Println(*conta)
}
