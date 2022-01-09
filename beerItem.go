package main

type BeerItem struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

type BeerItems struct {
	BeerItems []BeerItem `json:"beerItems"`
}
