package Migrations

import (
	"GoApi/internal/datalayer"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Migrator struct {
	context *sql.DB
}

var createTableAlbumsSQL string = `
	CREATE TABLE IF NOT EXISTS Albums (
		id SERIAL PRIMARY KEY,
		price FLOAT NOT NULL,
		name TEXT NOT NULL,
		artist TEXT NOT NULL
	);`

func NewMigrator(driver string, connection string) *Migrator {
	return &Migrator{
		context: datalayer.GetContext(driver, connection),
	}
}

func (m *Migrator) Migrate() {
	m.Execute(createTableAlbumsSQL)
}

func (m *Migrator) Execute(queries ...string) {
	for _, query := range queries {
		result, err := m.context.Query(query)
		log.Print(result, err)
	}
	m.context.Close()
}
