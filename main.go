package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const REQUESTS_PER_SECOND = 10

func main() {
	db := InitialiseDB()
	defer db.Close()

	r := gin.Default()

	r.Use(rateLimiter(REQUESTS_PER_SECOND))
	r.GET("/shorten", handleShorten(db))
	r.GET("/lengthen", handleLengthen(db))

	r.Run(":3000")
}

func rateLimiter(frequency int) gin.HandlerFunc {
	ticker := time.NewTicker(time.Second / time.Duration(frequency))
	channel := make(chan time.Time, 1)

	go func() {
		for t:= range ticker.C {
			channel <- t
		}
	}()

	return func(c *gin.Context) {
		select {
		case <-channel:
			c.Next()
		default:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded, try again"})
			c.Abort()
		}
	}
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