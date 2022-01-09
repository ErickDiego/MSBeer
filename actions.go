package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type BeerItem struct {
// 	Id       int64   `json:"id"`
// 	Name     string  `json:"name"`
// 	Brewery  string  `json:"brewery"`
// 	Country  string  `json:"country"`
// 	Currency string  `json:"currency"`
// 	Price    float64 `json:"price"`
// }

// type BeerItems struct {
// 	BeerItems []BeerItem `json:"beerItems"`
// }

var listBeerItems = []BeerItem{
	{
		Id:       0,
		Brewery:  "Autral",
		Country:  "Punta Arenas",
		Currency: "CLP",
		Name:     "Autral",
		Price:    1000},
	{
		Id:       1,
		Brewery:  "CCU",
		Country:  "Santiago",
		Currency: "CLP",
		Name:     "Cristal",
		Price:    500}}

type Request struct {
	Status  int         `json:"status"`
	Mensaje string      `json:"mensaje"`
	Data    interface{} `json:"data"`
}

// type BeerBox struct {
// 	PriceTotal float64 `json:"priceTotal"`
// }

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienbenido a MS Beer")
}

func GetListBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(listBeerItems)
}

func BeerAdd(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)

	var beerData BeerItem

	var beerItem BeerItem
	tieneDatos := false
	err := data.Decode(&beerData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	//Variable para almacenar datos de entrada
	beerItem = beerData

	beerData, tieneDatos = findBeer(beerData.Id)

	if !tieneDatos {
		listBeerItems = append(listBeerItems, beerItem)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "Se ha insertado Corectamente")

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		fmt.Fprintf(w, "El ID de la cerveza ya existe")
	}

}

// var requestValidator = govalidator.MapData{
// 	"id": []string{"required"},
// }

func GetBeer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// // rules := govalidator.MapData{
	// // 	"id": []string{"required"},
	// // }

	// validador := govalidator.Options{
	// 	Request: r,
	// 	Rules:   requestValidator,
	// 	Data:    params["id"]}

	// v := govalidator.New(validador)
	// err1 := v.ValidateJSON()

	// /*Retornara un error si se a detectado falla en la estructura*/
	// fmt.Println(err1)
	// fmt.Println("-> aqioooooo")

	// if len(err1) > 0 {
	// 	json.NewEncoder(w).Encode(
	// 		map[string]interface{}{
	// 			"errors": err1,
	// 		},
	// 	)
	// 	return
	// }

	Id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		fmt.Println(Id)
	}

	var beerItem BeerItem
	var request Request
	tieneDatos := false

	beerItem, tieneDatos = findBeer(Id)

	if tieneDatos {
		request.Status = 200
		request.Mensaje = "Se ha encontrado registros."
		request.Data = beerItem

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(request)
	} else {
		request.Status = 404
		request.Mensaje = "No ha encontrado registros."
		request.Data = beerItem

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(request)
	}

}

func GetBoxPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		fmt.Println(Id)
	}

	var beerItem BeerItem
	var request Request
	var beerBox BeerBox

	tieneDatos := false

	beerItem, tieneDatos = findBeer(Id)

	if tieneDatos {
		beerBox.PriceTotal = beerItem.Price * 6

		request.Status = 200
		request.Mensaje = "Operacion exitosa"
		request.Data = beerBox

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(request)

	} else {
		request.Status = 404
		request.Mensaje = "El Id de la cerveza no existe"
		request.Data = beerItem

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(request)
	}

}

//Funcion para buscar BeerItem y saber si existe en los registros
func findBeer(Id int64) (beer BeerItem, tiene_datos bool) {
	var datos BeerItem
	tieneDatos := false

	for _, item := range listBeerItems {

		if item.Id == Id {
			datos = item
			tieneDatos = true
			break
		}
	}

	return datos, tieneDatos
}
