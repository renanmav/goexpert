package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEPResponse
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		fmt.Println(data)

		// Create file if not exists
		var file *os.File
		if file, err = os.OpenFile("./cidade.txt", os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			if file, err = os.Create("./cidade.txt"); err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
			}
		}
		defer file.Close()

		if _, err = file.WriteString(fmt.Sprintf("%+v\n", data)); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
	}
}
