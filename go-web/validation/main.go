package main

/*
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro nuevo registro (sin tener una variable de último ID a nivel global).

Se debe implementar las validaciones de los campos al momento de enviar la petición, para eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400 con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).

Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se deben seguir los siguientes pasos::
1. Al momento de enviar la petición se debe validar que un token sea enviado 2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado). 3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un mensaje que “no tiene permisos para realizar la petición solicitada”.

*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Color     string  `json:"color" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Stock     uint    `json:"stock" validate:"required"`
	Code      string  `json:"code" validate:"required"`
	Published bool    `json:"published" validate:"required"`
}

var products = []*Product{
	{ID: 1, Name: "car", Color: "red", Price: 100, Stock: 4, Code: "c123", Published: true},
	{ID: 1, Name: "bicycle", Color: "red", Price: 100, Stock: 4, Code: "cb123", Published: true},
}

// use a single instance of Validate
var validate *validator.Validate

func main() {
	validate = validator.New()

	r := gin.Default()
	r.POST("/products", Create)
	r.Run()
}

func Create(c *gin.Context) {

	token := c.GetHeader("token")

	if token != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You do not have permission to perform the requested operation",
		})
		return
	}

	var request Product

	// TODO implements validation using validate library

	// err := validate.Struct(request)

	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fmt.Printf("The field %v is required", err.Field())
	// 	}
	// 	return
	// }

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.ID = len(products) + 1
	products = append(products, &request)
	c.JSON(http.StatusOK, request)

}
