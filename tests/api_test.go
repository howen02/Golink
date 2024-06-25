package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestShortenEndpoint(t *testing.T) {
	fmt.Println("Testing shorten endpoint")
	resp, err := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

	if err != nil || resp.StatusCode != http.StatusOK{
		log.Fatal("Error getting shortUrl: ", err)
	}	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	fmt.Println(string(body))
	fmt.Println("")
}

func TestLengthenEndopint(t *testing.T) {
	fmt.Println("Testing lengthen endpoint")
	resp, err := http.Get("http://localhost:3000/lengthen?shortUrl=JT0UJwME")

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Error getting longUrl: ", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	fmt.Println(string(body))
	fmt.Println("")
}

func TestRateLimiterExceed(t *testing.T) {
	requests := 20
	fmt.Println("Testing rate limiter with exceeding requests")

	for i := 0; i < requests; i++ {
		resp, err := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

		if err != nil && resp.StatusCode != http.StatusTooManyRequests {
			log.Fatal("Rate limiter test failed")
		}
	}

	fmt.Printf("Rate limiter test passed with %d requests\n", requests)
	fmt.Println("")	
}

func TestRateLimiterWithin(t *testing.T) {
	requests := 10
	fmt.Println("Testing rate limiter within requests")

	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:3000/shorten?longUrl=http://www.google.com")

		if err != nil && resp.StatusCode != http.StatusOK {
			log.Fatal("Rate limiter test failed")
		}
	}

	fmt.Printf("Rate limiter test passed with %d requests\n", requests)
	fmt.Println("")
}