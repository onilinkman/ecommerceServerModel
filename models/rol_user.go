package models

import "main/databasee"

type Rol_user struct {
	Id_rol  int `json:"id"`
	Id_user int `json:"id_user"`
}

type Rols_user []Rol_user

const schemaRol_user string = `CREATE TABLE if NOT EXISTS rol_user(
	id_rol int NOT NULL,
	id_user int NOT NULL,
	CONSTRAINT fk_rol
	FOREIGN KEY (id_rol) REFERENCES rol(id_rol) ON DELETE CASCADE,
	CONSTRAINT fk_user
	FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE)`

func CreateTableRol_user() {
	databasee.ExecuteExec(schemaRol_user)
}
