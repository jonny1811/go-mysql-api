package connect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Funcion para conectarnos a la DB
var Db *sql.DB

func Connect() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+
		"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}
	Db = connection

}

// Cerrar la Conexión
func CloseConnection() {
	Db.Close()
}
