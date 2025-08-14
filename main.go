package main

import (
	"clase_3_mysql_driver/handlers"
	"clase_3_mysql_driver/models"
)

func main() {
	// handlers.List()
	// handlers.ListById(1)
	cliente := models.Cliente{Nombre: "Marcos Ñamandu", Correo: "marcos@hotmail.com",
		Telefono: "+96656898"}
	handlers.Insert(cliente)
	handlers.List()
}
