package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CurrencyQuote struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300*time.Millisecond))
	defer cancel()

	endpoint, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	request, err := http.DefaultClient.Do(endpoint)
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
