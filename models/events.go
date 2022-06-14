package models

import (
	"main/databasee"
)

/* CURRENT_TIMESTAMP + INTERVAL 1 DAY */
const scheduleClear_Session string = `CREATE EVENT IF NOT EXISTS clear_sessions
								ON SCHEDULE EVERY 1 DAY
								STARTS '2022-01-19 18:00:00'
								DO
									DELETE FROM sessions 
									WHERE date_due <= CURRENT_TIMESTAMP`

const globalEvent_Schedule string = `SET GLOBAL event_scheduler = ON`

func SetGlobalEvent() {
	databasee.ExecuteExec(globalEvent_Schedule)
}

func CreateEventClearSession() {
	databasee.ExecuteExec(scheduleClear_Session)
}
