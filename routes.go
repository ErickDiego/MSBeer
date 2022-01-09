package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	name       string
	method     string
	pattern    string
	handleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		/*Añadir los datos al router*/
		router.
			Name(route.name).
			Methods(route.method).
			Path(route.pattern).
			Handler(route.handleFunc)
	}
	return router
}

/*
	arreglo de rutas para accer a los metodos a travez de Rest
*/
var routes = []Route{
	{"Index", "GET", "/", Index},
	{"GetListBeers", "GET", "/getListBeers/", GetListBeers},
	{"BeerAdd", "POST", "/BeerAdd", BeerAdd},
	{"GetBeer", "GET", "/GetBeer/{id}", GetBeer},
	{"GetBoxPrice", "GET", "/beers/{id}/boxprice", GetBoxPrice}}
