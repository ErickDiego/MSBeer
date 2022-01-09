# MSBeer

## Introduccion

Bender es fanático de las cervezas, y quieretener un registro de todas las cervezas queprueba y cómo calcular el precio que necesitapara comprar una caja de algún tipoespecifico de cervezas. Para esto necesitauna API REST con esta información que posteriormente compartir con sus amigos.

## Descripcion

Se solicita crear un API REST basándonos en la definición que se encuentra en el archivo: <https://bitbucket.org/lgaetecl/microservices-test/src/master/openapi.yaml/>

## Funcionalidad

- Listado de Cervezas
    Metodo: _GET_
    Path: <http://localhost:8080/getListBeers/>

- Insercion de Cervezas
    Metodo: _POST_
    Path: <http://localhost:8080/BeerAdd>
    Body:

    ```
      {
        "Id": 3,
        "Brewery": "CCU",
        "Country": "Santiago",
        "Currency": "CLP",
        "Name": "Dorada",
        "Price": 500
        }
    ```

- Busqueda de Cerveza
    Metodo: GET
    Path: <http://localhost:8080/GetBeer/{Id>}
    Header: Id = 1

- Calculo valor de caja de cerveza
    Metodo: GET
    Path: <http://localhost:8080/beers/{id}/boxprice>
    Header: Id = 1

## Construido con 🛠️

- [Gorila Mux](https://github.com/gorilla/mux) - Libreria para crear API Rest

## Comandos

- Ejecucion del proyecto de manera local

    ```
    go run .
    ```

- Detencion del servicio

    ```
    ctrl + c 
    ```