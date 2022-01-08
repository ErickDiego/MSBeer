package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var listBeer = []Beer{
	{Name: "Patagonia", Price: 3000},
	{Name: "Escudo", Price: 3000},
	{Name: "Bear Beer", Price: 3000},
	{Name: "Autral", Price: 3000}}

/*Sacar el Struct a una clase nueva*/
type Beer struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Beers struct {
	Beers []Beer `json:"beers"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor Web con GO")
}

func getListBeers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(listBeer)
}

func insertBeer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor Web con GO")
}
