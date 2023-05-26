package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/lclpedro/goexpert-desafios/pkg/requester"
)

var DB *sqlx.DB

func main() {
	mux := http.NewServeMux()
	DB = NewDatabase()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Serviço de cotação de Dólar"}`))
	})

	mux.HandleFunc("/cotacao", GetQuote)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

type CurrencyQuote struct {
	Currency struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetQuote")
	defer fmt.Println("GetQuote Finish")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(200*time.Millisecond))
	defer cancel()

	client := requester.NewRequester(ctx)
	request, err := client.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error in request quote"}`))
		return
	}

	response, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error in request quote"}`))
		return
	}

	var currencyQuote CurrencyQuote

	err = json.Unmarshal(response, &currencyQuote)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error in request quote"}`))
		return
	}

	apiResponse := QuotesResponse{
		Code:       currencyQuote.Currency.Code,
		Codein:     currencyQuote.Currency.Codein,
		Name:       currencyQuote.Currency.Name,
		High:       currencyQuote.Currency.High,
		Low:        currencyQuote.Currency.Low,
		Bid:        currencyQuote.Currency.Bid,
		CreateDate: currencyQuote.Currency.CreateDate,
	}

	err = SaveQuote(apiResponse)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error in request quote"}`))
		return
	}

	data, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error in request quote"}`))
		return
	}
	w.Write(data)
}

func NewDatabase() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "../database.db")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	initDatabase := `
	CREATE TABLE IF NOT EXISTS quotes (
    code TEXT,
    codein TEXT,
    name TEXT,
    high TEXT,
    low TEXT,
    bid TEXT,
		create_date TEXT
	);`
	db.MustExec(initDatabase)
	return db
}

type QuotesResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	Bid        string `json:"bid"`
	CreateDate string `json:"create_date"`
}

type QuotesModel struct {
	Code       string `db:"code"`
	Codein     string `db:"codein"`
	Name       string `db:"name"`
	High       string `db:"high"`
	Low        string `db:"low"`
	Bid        string `db:"bid"`
	CreateDate string `db:"create_date"`
}

func SaveQuote(quote QuotesResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Millisecond))
	defer cancel()

	toSave := &QuotesModel{
		Code:       quote.Code,
		Codein:     quote.Codein,
		Name:       quote.Name,
		High:       quote.High,
		Low:        quote.Low,
		Bid:        quote.Bid,
		CreateDate: quote.CreateDate,
	}

	_, err := DB.NamedExecContext(
		ctx,
		`INSERT INTO quotes (
			code, codein, name, high, low, bid, create_date
			) VALUES (:code, :codein, :name, :high, :low, :bid, :create_date)`,
		toSave,
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
