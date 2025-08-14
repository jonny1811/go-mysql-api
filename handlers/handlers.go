package handlers

import (
	"bufio"
	"clase_3_mysql_driver/connect"
	"clase_3_mysql_driver/models"
	"fmt"
	"log"
	"os"
	"strconv"
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

func Edit(cliente models.Cliente, id int) {
	connect.Connect()
	sql := "UPDATE clientes set nombre=?, correo=?, telefono=? WHERE id=?;"
	result, err := connect.Db.Query(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se editó el registro exitosamente")
}

func Delete(id int) {
	connect.Connect()
	sql := "DELETE FROM clientes WHERE id=?;"
	_, err := connect.Db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se eliminó el registro exitosamente")
}

// ##################FUNCIONES DE TRABAJO
var ID int
var nombre, correo, telefono string

func Execute() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opción : ")
	fmt.Println("1- Listar clientes")
	fmt.Println("2- Listar cliente por ID")
	fmt.Println("3- Crear cliente")
	fmt.Println("4- Editar cliente")
	fmt.Println("5- Eliminar cliente")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				List()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				ListById(ID)
				return
			}
			if scanner.Text() == "3" {
				fmt.Println("Ingrese el Nombre : ")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("Ingrese el Email : ")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese el Telefono : ")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insert(cliente)
				return
			}
			if scanner.Text() == "4" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				fmt.Println("Ingrese el Nombre : ")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("Ingrese el Email : ")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese el Telefono : ")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insert(cliente)
				Edit(cliente, ID)
				return
			}
			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				Delete(ID)
				return
			}
		}
	}
}
