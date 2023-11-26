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
	Nome  string
	Ativo bool
}

func (e Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %s foi desativada", e.Nome)
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{
		Nome:  "Full Cycle",
		Ativo: true,
	}

	Desativacao(minhaEmpresa)

	fmt.Printf("%+v", wesley)
}
