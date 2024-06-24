package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := InitialiseDB()
	defer db.Close()

	r := gin.Default()

	r.GET("/shorten", func(c *gin.Context) {
		longUrl := c.Query("longUrl")

		if longUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "longUrl query parameter is required"})
			return
		}

		shortUrl := InsertLongUrl(db, longUrl)

		if shortUrl == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "shortUrl does not exist in database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"shortUrl": shortUrl})
	})

	r.GET("/lengthen", func(c *gin.Context) {
		shortUrl := c.Query("shortUrl")

		if shortUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl query parameter is required"})
			return
		}

		longUrl := GetLongUrl(db, shortUrl)

		if longUrl == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating longUrl"})
			return
		}


		log.Println("Redirecting to: ", longUrl)
		c.Redirect(http.StatusMovedPermanently, longUrl)
	})

	r.Run(":3000")
}