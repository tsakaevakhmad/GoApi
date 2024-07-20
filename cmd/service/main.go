package main

import (
	"GoApi/internal/businesslayer/services"
	"GoApi/internal/controllers"
	"GoApi/internal/datalayer/repositories"
)

func main() {
	album := controllers.NewAlbumController(services.NewAlbumServices(repositories.NewAlbumRepository()))
	Run(album)
}
