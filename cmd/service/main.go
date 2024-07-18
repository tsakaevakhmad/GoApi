package main

import (
	"GoApi/internal/businesslayer/services"
	"GoApi/internal/controllers"
)

func main() {
	album := controllers.NewAlbumController(services.NewAlbumServices())
	InitControllers(album).Run("localhost:8080")
}
