package databasee

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //driver de la base de datos
)

var db *sql.DB

const username string = "root"
const password string = "pelota12"
const host string = "192.168.0.128"
const port int = 3306
const database string = "ecommerce"

func init() {
	CreateConnection()
}

func CreateConnection() {
	connection, err := sql.Open("mysql", generateURL())
	if err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexión a la base de datos establecida")
	}
}

func ExecuteExec(schema string, args ...interface{}) {

	_, err := db.Exec(schema, args...)
	if err != nil {
		fmt.Println(err)
	}
}

func ExecuteQuery(schema string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(schema, args...)

	return rows, err
}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func CloseConnection() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
}
