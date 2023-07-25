package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/gin-contrib/cors"
)

func main() {
    port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default)
}