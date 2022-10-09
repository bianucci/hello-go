package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = map[string]album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, maps.Values(albums))
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	if album, ok := albums[id]; ok {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("Album not found for id=\"%s\"", id))
}

func main() {
	router := gin.Default()

	public := router.Group("/api")

	public.GET("/albums", getAlbums)
	public.GET("/albums/:id", getAlbum)
	public.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
