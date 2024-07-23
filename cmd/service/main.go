package main

import (
	"GoApi/internal/businesslayer/services"
	"GoApi/internal/controllers"
	"GoApi/internal/datalayer/Migrations"
	"GoApi/internal/datalayer/repositories"
)

func main() {
	dbConnection := "user=postgres password=password dbname=postgres sslmode=disable"
	driver := "postgres"
	Migrations.NewMigrator(driver, dbConnection).Migrate()
	album := controllers.NewAlbumController(services.NewAlbumServices(repositories.NewAlbumRepository(driver, dbConnection)))
	Run(album)
}
