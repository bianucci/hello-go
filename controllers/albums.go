package controllers

import (
	"fmt"
	"net/http"

	"github.com/bianucci/hello-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

var albums = map[string]models.Album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, maps.Values(albums))
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	if album, ok := albums[id]; ok {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("Album not found for id=\"%s\"", id))
}
