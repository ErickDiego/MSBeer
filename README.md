# MSBeer

# Introduccion
    Bender es fanático de las cervezas, y quiere tener un registro de todas las cervezas que prueba y cómo calcular el precio que necesita para comprar una caja de algún tipo especifico de cervezas. Para esto necesita una API REST con esta información que posteriormente compartir con sus amigos.

# Descripcion
    Se solicita crear un API REST basándonos en la definición que se encuentra en el archivo: https://bitbucket.org/lgaetecl/microservices-test/src/master/openapi.yaml

# Funcionalidad
- Listado de Cervezas
    http://localhost:8080/getListBeers/

- Insercion de Cervezas 
    http://localhost:8080/BeerAdd

    Json de Prueba
    ```
      {
        "Id": 3,
        "Brewery": "CCU",
        "Country": "Santiago",
        "Currency": "Peso Chileno",
        "Name": "Dorada",
        "Price": 500
        }
    ```

- Busqueda de Cerveza
   http://localhost:8080/GetBeer/1
   
- Calculo valor de caja de cerveza
    http://localhost:8080/beers/{id}/boxprice


## Construido con
[Gorila Mux] (github.com/gorilla/mux)