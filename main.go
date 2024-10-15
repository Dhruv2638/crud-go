package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Nayab", Artist: "seedhe maut", Price: 45.98},
	{ID: "2", Title: "Hard drive", Artist: "raftaar", Price: 90.98},
	{ID: "3", Title: "Flying towardes the city", Artist: "panther", Price: 60.98},
}

// get all the Albums
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// Add new Album to data
func addAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// Get the Album by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	//  get all the albums
	router.Handle("GET", "/albums", getAlbums)
	router.Handle("POST", "/album", addAlbum)
	router.Handle("GET", "/albums/:id", getAlbumById)
	router.Run("localhost:8080")
}
