package datalayer

import "database/sql"
import _ "github.com/lib/pq"

func GetContext(driver string, connection string) *sql.DB {
	context, err := sql.Open(driver, connection)
	if err != nil {
		panic(err)
	}
	return context
}
