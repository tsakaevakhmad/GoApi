package controllers

import (
	"GoApi/internal/businesslayer/interfaces"
	"GoApi/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumController struct {
	albumServices interfaces.IAlbumServices
}

func NewAlbumController(albumServices interfaces.IAlbumServices) *AlbumController {
	return &AlbumController{
		albumServices: albumServices,
	}
}

func (a *AlbumController) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, a.albumServices.GetAll())
}

func (a *AlbumController) GetAlbumsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	album := a.albumServices.Get(id)

	if album == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found 404"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (a *AlbumController) PostAlbum(c *gin.Context) {
	var album domain.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	a.albumServices.Create(album)
	c.IndentedJSON(http.StatusOK, album)
}

func (a *AlbumController) PutAlbum(c *gin.Context) {
	var album domain.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	a.albumServices.Put(album)
	c.IndentedJSON(http.StatusOK, album)
}

func (a *AlbumController) DeleteAlbum(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if nil == a.albumServices.Delete(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found 404"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}
