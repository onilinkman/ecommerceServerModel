package models

import "main/databasee"

//Rol defined the structure of the object rol
type Rol struct {
	Id_rol      int    `json:"id"`
	Description string `json:"description"`
}

//Rols defined the structure of the object rols
type Rols []Rol

const schemaRol string = `CREATE TABLE if NOT EXISTS rol(
	id_rol INT PRIMARY KEY NOT NULL,
	description varchar(50) NOT NULL)`

func CreateTableRol() {
	databasee.ExecuteExec(schemaRol)
}
