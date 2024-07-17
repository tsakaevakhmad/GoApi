package controllers

import (
	"GoApi/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var albums = []domain.Album{
	{Id: 1, Price: 13.99, Artist: "Oleg F. A.", Name: "Her'"},
	{Id: 2, Price: 10.99, Artist: "Manis K. K.", Name: "YouNga"},
	{Id: 3, Price: 9.99, Artist: "Alister K. E.", Name: "Shorts"},
	{Id: 4, Price: 2.99, Artist: "Capitan K. E.", Name: "Manual"},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found 404"})
}

func PostAlbum(c *gin.Context) {
	var album domain.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	album.Id = int64(len(albums) + 1)
	albums = append(albums, album)
	c.IndentedJSON(http.StatusOK, album)
}

func PutAlbum(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var album domain.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	if id != album.Id {
		c.IndentedJSON(http.StatusBadRequest, album)
		return
	}

	for _, a := range albums {
		if a.Id == id {
			a = album
			c.IndentedJSON(http.StatusOK, album)
		}
	}

	albums = append(albums, album)
}
