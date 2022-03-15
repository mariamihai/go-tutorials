package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var initAlbums = []album{
	{ID: "1", Title: "Title 1", Artist: "Artist 1", Price: 10.12},
	{ID: "2", Title: "Title 2", Artist: "Artist 2", Price: 14.48},
	{ID: "3", Title: "Title 3", Artist: "Artist 1", Price: 9.56},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	// can be replaced with Context.JSON to send a more compact JSON
	c.IndentedJSON(http.StatusOK, initAlbums)
}
