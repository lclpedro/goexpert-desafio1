package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lclpedro/goexpert-desafios/desafio1/pkg/requester"
)

type CurrencyQuote struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(300*time.Millisecond))
	defer cancel()

	client := requester.NewRequester(ctx)
	request, err := client.Get("http://127.0.0.1:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)

	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}

	var data CurrencyQuote
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../cotacoes.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte(fmt.Sprintf("Dólar: %s", data.Bid)))

	fmt.Println("Finish")
}
