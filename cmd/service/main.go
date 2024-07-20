package main

import (
	"GoApi/internal/businesslayer/services"
	"GoApi/internal/controllers"
	"GoApi/internal/datalayer/repositories"
)

func main() {
	dbConnection := "user=postgres password=password dbname=AbumsPG sslmode=disable"
	driver := "postgres"
	album := controllers.NewAlbumController(services.NewAlbumServices(repositories.NewAlbumRepository(driver, dbConnection)))
	Run(album)
}
