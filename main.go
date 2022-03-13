package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbums)
	router.PUT("/albums/:id", updateAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for k, a := range albums {
		if a.ID == id {
			albums = append(albums[:k], albums[k+1:]...)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
	for k := range albums {
		if albums[k].ID == id {
			err := c.BindJSON(&albums[k])
			if err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, albums[k])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
