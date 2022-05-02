package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var albums = sync.Map{}
var artists = sync.Map{}

var _albums = map[string]album{
	"1": {ID: "1", Title: "Ænima", Artist: "Tool", Price: 56.99},
	"2": {ID: "2", Title: "Amnesiac", Artist: "Radiohead", Price: 17.99},
}

var _artists = map[string]artist{
	"1": {ID: "1", Name: "Tool"},
	"2": {ID: "2", Name: "Radiohead"},
	"3": {ID: "3", Name: "Aphex Twin"},
}

func seed() {
	for _, x := range _albums {
		albums.Store(x.ID, x)
	}
	for _, x := range _artists {
		artists.Store(x.ID, x)
	}
}

func main() {
	seed()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/artists", getArtists)
	router.POST("/artists", postArtists)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbum)
	router.DELETE("/artists/:id", deleteArtist)
	router.GET("/artists/:id", getArtistByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.PUT("/artists/:id", updateArtist)
	router.Run("0.0.0.0:8080")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type artist struct {
	ID string `json:"id"`
	Name string `json:"name"`
}



func getAlbums(c *gin.Context) {
	a := []any{}
	albums.Range(func (_ any, value any) bool {
		a = append(a, value)
		return true
	})

	c.IndentedJSON(http.StatusOK, a)
}

func getArtists(c *gin.Context) {
	a := []any{}
	artists.Range(func (_ any, value any) bool {
		a = append(a, value)
		return true
	})

	c.IndentedJSON(http.StatusOK, a)
}

func postArtists(c *gin.Context) {
	var artist artist

	if err := c.BindJSON(&artist); err != nil {
		return
	}

	artists.Store(artist.ID, artist)

	c.IndentedJSON(http.StatusCreated, artist)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums.Store(newAlbum.ID, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, ok := albums.Load(id)
	if ok {
		c.IndentedJSON(http.StatusOK, album)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func getArtistByID(c *gin.Context) {
	id := c.Param("id")

	artist, ok := artists.Load(id)
	if ok {
		c.IndentedJSON(http.StatusOK, artist)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	a, ok := albums.LoadAndDelete(id)
	if ok {
		c.IndentedJSON(http.StatusOK, a)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func deleteArtist(c *gin.Context) {
	id := c.Param("id")

	a, ok := artists.LoadAndDelete(id)
	if ok {
		c.IndentedJSON(http.StatusOK, a)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	_, ok := albums.Load(id)
	if ok {
		var newAlbum album

		if err := c.BindJSON(&newAlbum); err != nil {
			c.AbortWithStatus(403)
			return
		}

		albums.Store(id, newAlbum)
		c.IndentedJSON(http.StatusOK, newAlbum)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func updateArtist(c *gin.Context) {
	id := c.Param("id")

	_, ok := artists.Load(id)
	if ok {
		var artist artist

		if err := c.BindJSON(&artist); err != nil {
			c.AbortWithStatus(403)
			return
		}
		artists.Store(id, artist)
		c.IndentedJSON(http.StatusOK, artist)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
}