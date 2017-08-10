package main

import (
	"bytes"
    "log"
    "net/http"
    "os"
    "strconv"
    
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	_ "github.com/heroku/x/hmetrics/onload"
)

var (
	repeat int
)

const SERVER_PORT string = "8080"

func repeatHandler(c *gin.Context) {
    var buffer bytes.Buffer
    for i := 0; i < repeat; i++ {
        buffer.WriteString("Hello from Go!\n")
    }
    c.String(http.StatusOK, buffer.String())
}

func main() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		port = SERVER_PORT
	}

	tStr := os.Getenv("REPEAT")
    repeat, err = strconv.Atoi(tStr)
    if err != nil {
        log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
        repeat = 5
    }

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/mark", func(c *gin.Context) {
  		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})
	
	router.GET("/repeat", repeatHandler)

	router.Run(":" + port)
}
