package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func GetBeer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

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

	ConsumirAPI()

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

type Currency struct {
	Success   string      `json:"success"`
	Terms     string      `json:"terms"`
	Privacy   string      `json:"privacy"`
	Timestamp int         `json:"timestamp"`
	Source    string      `json:"source"`
	Quotes    interface{} `json:"quotes"`
}


func ConsumirAPI() {
	fmt.Println("Calling API...")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.currencylayer.com/live?access_key=03d6c422fccef07abae7c5aae58b4088&currencies=CLP", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Currency
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
