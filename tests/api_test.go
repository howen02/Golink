package main

import (
	"log"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestShortenEndpoint(t *testing.T) {
	log.Println("Testing shorten endpoint")
	resp, err := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

	if err != nil || resp.StatusCode != http.StatusOK{
		log.Fatal("Error getting shortUrl: ", err)
	}	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	log.Println(string(body))
	log.Println("")
}

func TestLengthenEndopint(t *testing.T) {
	log.Println("Testing lengthen endpoint")
	resp, err := http.Get("http://localhost:3000/lengthen?shortUrl=JT0UJwME")

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Error getting longUrl: ", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	log.Println(string(body))
	log.Println("")
}

func TestRateLimiterExceed(t *testing.T) {
	time.Sleep(time.Second)
	requests := 20
	tooManyRequests := false
	log.Println("Testing rate limiter with exceeding requests")

	for i := 0; i < requests; i++ {
		resp, _ := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

		if resp.StatusCode == http.StatusTooManyRequests {
			tooManyRequests = true
			break
		}
	}

	if (!tooManyRequests) {
		log.Fatal("Rate limiter test failed")
	}

	log.Printf("Rate limiter test passed with %d requests\n", requests)
	log.Println("")	
}

func TestRateLimiterWithin(t *testing.T) {
	time.Sleep(time.Second)
	requests := 10
	log.Println("Testing rate limiter within requests")

	for i := 0; i < 10; i++ {
		resp, _ := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

		if resp.StatusCode != http.StatusOK {
			log.Fatal("Rate limiter test failed")
		}

		time.Sleep(100 * time.Millisecond)
	}

	log.Printf("Rate limiter test passed with %d requests\n", requests)
	log.Println("")
}
