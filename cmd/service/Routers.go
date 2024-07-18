package main

import (
	"GoApi/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitControllers(album *controllers.AlbumController) *gin.Engine {
	router := gin.Default()

	// Albums
	router.GET("/album", album.GetAlbums)
	router.GET("/album/:id", album.GetAlbumsById)
	router.POST("/album/create", album.PostAlbum)
	router.PUT("/album/put", album.PutAlbum)
	router.DELETE("/album/delete/:id", album.DeleteAlbum)

	return router
}
