package main

import (
	"GoApi/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitControllers() *gin.Engine {
	router := gin.Default()

	router.GET("/album/all", controllers.GetAlbums)
	router.GET("/album/:id", controllers.GetAlbumsById)
	router.POST("/album/create", controllers.PostAlbum)
	router.PUT("/album/put", controllers.PutAlbum)

	return router
}
