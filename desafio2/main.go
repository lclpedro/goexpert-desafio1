package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/lclpedro/goexpert-desafios/pkg/requester"
)


// Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.
// As duas requisições serão feitas simultaneamente para as seguintes APIs:
// https://cdn.apicep.com/file/apicep/" + cep + ".json
// http://viacep.com.br/ws/" + cep + "/json/
// Os requisitos para este desafio são:
// - Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
// - O resultado da request deverá ser exibido no command line, bem como qual API a enviou.
// - Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

func clientViaCEP(ctx context.Context, cep string, result chan string, apiClient chan string) {
	client := requester.NewRequester(ctx)
	request, err := client.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	result <- string(body)
	apiClient <- "VIA CEP"
}

func clientAPICep(ctx context.Context, cep string, result chan string, apiClient chan string) {
	client := requester.NewRequester(ctx)
	request, err := client.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	result <- string(body)
	apiClient <- "API CEP"
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	var result = make(chan string)
	var client = make(chan string)

	go clientViaCEP(ctx, "01001-000", result, client)
	go clientAPICep(ctx, "01001000", result, client)

	endereco := <-result
	apiClient := <-client

	fmt.Println(apiClient)
	fmt.Println(endereco)
}
