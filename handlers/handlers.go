package handlers

import (
	"clase_3_mysql_driver/connect"
	"clase_3_mysql_driver/models"
	"fmt"
	"log"
)

func Listar() {
	connect.Connect()
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"
	datos, err := connect.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer connect.CloseConnection()
	/*
		clientes := models.Clientes{}
		for datos.Next() {
			dato := models.Cliente{}
			datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
			clientes = append(clientes, dato)
		}
		fmt.Println(clientes)*/
	fmt.Println("\nLISTADO DE CLIENTES")
	fmt.Println("---------------------------------------------------------")
	for datos.Next() {
		var dato models.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | Email: %s | Teléfono: %s \n", dato.Id, dato.Nombre,
			dato.Correo, dato.Telefono)
		fmt.Println("---------------------------------------------------------")
	}
}
