# golink
#### Video Demo:  <URL HERE>
#### Description:
golink is a simple URL shortener API built with the programming language go. 

## Packages
Here are some of the packages used:
- database/sql: To interact with a SQL database to store and retrieve URLs
- net/http: To provide different http status codes returned when an API request is made
- time: Used in rate limiting and providing delay between tests to prevent our tests from triggering the rate limiter
- gin-gonic/gin: To create a http server, provide routing, middleware and also context management
- log: To log the different processes occuring in the API for observation
- io: To read the response body of http requests in the test file
- testing: To write unit tests
- crypto/sha256: To hash long URLs into a short code
- encoding/base64: To encode the hashed URL into a safe text string
- mattn/go-sqlite3: To access database drivers for SQLite3

## Structure
Let's have a look at how the project is structured
```
golink/
├── bin/
│   └── main
├── tests/
│   └── api_tests.go
├── main.go
├── golink.db
├── Makefile
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```
## Makefile
I utilised a makefile to make the development process smoother, here are some of the commands I used throughout this project
- build: This command compiles the entire go program and outputs the binary as `main` in the `bin` directory
- run: This command runs the build command and then executes the binary 
- clean: This command clears the binary
- test: This command runs `go test` to run all my testing files

## Initialisation
We first initialise the db by creating a table for Urls if they do not exist yet. For the schema design we have a simple table with three columbs being id which is an auto-incrementing idendification number for our URLs. THen we have longUrl and shortUrl which are text types that cannot be null.

## Endpoints
It features five different endpoints along with features such as rate limiting and testing. Let's have a look at the different endpoints:

### /shorten
This is the main endpoint where users can pass in a URL and obtain a shortened version of it instead. Here's how it works:
1. The API creates a hash to represent the URL
2. We insert the URL along with its hash into the database
3. Any error thrown will be handled
4. The hash will be returned to the user

With the given hash, the user can obtain the original URL by simply lengthening it

### /lengthen
This is where users can pass in hashed URLs to obtain the original version
1. The API finds a row with the hashed URL
2. It then returns the long URL

### /group/shorten & /group/lengthen
These endpoints are similar to the ones above, just that the user can pass in multiple URLs at a time instead of just one

### health
This endpoint pings the database to check the health status of it, notifying the user if there is a connection issue with the database

## Rate Limiting
I decided to add rate limiting to my API to mimic a realistic one and I had to implement go routines for this. 

How it works is that a ticker is created and it ticks at a fixed frequency. Say we set the `REQUESTS_PER_SECOND` to 10, then there will be 10 requests allowed in a second. The ticker then fires a response every 100 miliseconds which will be received by a channel, allowing an API request to be made.

If no response is received by the channel, it simply means that 100 miliseconds has not passed yet and any API requests will be denied.

## Testing
I decided to make a test file for this API as well as go has been providing a relatively nice developer experience so far. We have four different functions, testing for lengthening, shortening and both scenarios where we exceed the rate limit and stay within it. 

Testing is really as simple as running `go test` and the program simply runs all functions with the appropriate formatting in the folder
