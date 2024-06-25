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

	if err != nil || resp.StatusCode != http.StatusOK {
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