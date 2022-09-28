package app

import (
	"database/sql"
	"secret-service/lib"
	"time"
)

func CreatedTimeCompare(dbClinet *sql.DB) {
	for {
		time.Sleep(10 * time.Second)
		timeStr := time.Now().Add(-lib.TimeWhenDelete).Format(lib.DbTLayout)
		dbClinet.Exec("DELETE FROM users where created_at <= $1", timeStr)
	}
}
