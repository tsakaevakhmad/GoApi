package main

import (
	"GoApi/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Run(album *controllers.AlbumController) {
	router := gin.Default()

	// Albums
	router.GET("/album", album.GetAlbums)
	router.GET("/album/:id", album.GetAlbumsById)
	router.POST("/album/create", album.PostAlbum)
	router.PUT("/album/put", album.PutAlbum)
	router.DELETE("/album/delete/:id", album.DeleteAlbum)

	router.Run("localhost:8080")
}
