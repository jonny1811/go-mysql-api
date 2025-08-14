package main

import (
	"clase_3_mysql_driver/handlers"
	"clase_3_mysql_driver/models"
)

func main() {
	// handlers.List()
	// handlers.ListById(1)
	cliente := models.Cliente{Nombre: "Jose Amarilla", Correo: "jose.amarilla@hotmail.com",
		Telefono: "+9998653"}
	// handlers.Insert(cliente)
	handlers.Edit(cliente, 3)
	handlers.List()
}
