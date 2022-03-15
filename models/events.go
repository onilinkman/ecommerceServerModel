package models

import (
	"main/databasee"
)

const scheduleClear_Session string = `CREATE EVENT IF NOT EXISTS clear_sessions
								ON SCHEDULE EVERY 1 DAY
								STARTS CURRENT_TIMESTAMP + INTERVAL 1 DAY
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
