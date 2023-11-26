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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua dos bobos",
			Numero:     0,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	fmt.Printf("Name: %s, Idade %d, Ativo: %t, Endereço: %+v\n", wesley.Nome, wesley.Idade, wesley.Ativo, wesley.Endereco)

	wesley.Ativo = false
	wesley.Desativar()
}
