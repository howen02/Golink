package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := InitialiseDB()
	defer db.Close()

	r := gin.Default()

	r.GET("/shorten", handleShorten(db))
	r.GET("/lengthen", handleLengthen(db))

	r.Run(":3000")
}

func handleShorten(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func handleLengthen(db *sql.DB) gin.HandlerFunc {
	return func(c * gin.Context) {
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


		c.JSON(http.StatusOK, gin.H{"longUrl": longUrl})
	}
}