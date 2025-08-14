package handlers

import (
	"clase_3_mysql_driver/connect"
	"clase_3_mysql_driver/models"
	"fmt"
	"log"
)

func List() {
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

func ListById(id int) {
	connect.Connect()
	sql := "SELECT id, nombre, correo, telefono FROM clientes where id=?;"
	datos, err := connect.Db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer connect.CloseConnection()
	fmt.Println("\nLISTAR CLIENTE POR ID")
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

func Insert(cliente models.Cliente) {
	connect.Connect()
	sql := "INSERT INTO clientes values(null, ?, ?, ?);"
	result, err := connect.Db.Query(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se creó el registro exitosamente")
}
