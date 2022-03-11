package models

import (
	"main/databasee"
	"time"
)

type SessionUser struct {
	Id_user    int       `json:"id"`
	Id_session string    `json:"id_session"`
	Date       time.Time `json:"date"`
}

type SessionsUser []SessionUser

const UsersSessions string = `CREATE TABLE if NOT EXISTS sessions(
	id_session varchar(60) PRIMARY KEY NOT NULL,
	id_user INT NOT NULL,
	date_due TIMESTAMP,
	CONSTRAINT fk_user_session 
	FOREIGN KEY (id_user) 
	REFERENCES users(id) ON DELETE CASCADE)`

func CreateTableSession() {
	databasee.ExecuteExec(UsersSessions)
}

func InsertSession(id_user int, id_session string, date_due time.Time) {
	databasee.ExecuteExec(`INSERT INTO sessions(id_session,id_user,date_due) VALUES(?,?,?)`, id_session, id_user, date_due)
}

func GetSessionsByUserId(id int) (*SessionsUser, error) {
	return getSessionsByQuery(`SELECT id_session,id_user,date_due FROM sessions WHERE id_user=?`, id)
}

func DeleteSession(cookie string) {
	databasee.ExecuteExec(`DELETE FROM sessions WHERE id_session=?`, cookie)
}

func DeleteUserSession(cookie string) {
	databasee.ExecuteExec(`DELETE FROM sessions WHERE id_session=?`, cookie)
}

func getSessionsByQuery(query string, id int) (*SessionsUser, error) {
	rows, err := databasee.ExecuteQuery(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sessions = SessionsUser{}
	for rows.Next() {
		session := &SessionUser{}
		err = rows.Scan(&session.Id_session, &session.Id_user)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, *session)
	}
	return &sessions, nil
}
