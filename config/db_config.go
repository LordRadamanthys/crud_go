package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const db_credentials = "root:bebop268170@/cursogo"

func CreateConection(table string) (*sql.DB, error) {
	db, err := sql.Open("mysql", db_credentials)
	if err != nil {
		return nil, err
	}
	db.Exec(fmt.Sprintf("user %s", table))
	return db, nil
}
