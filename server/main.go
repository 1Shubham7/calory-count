package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/1shubham7/calory-count/server/routes"
)

func main() {
    port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default)

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", router.EntryById)
	router.GET("/ingredient/:ingredients", routes.GetEntriesByIngredient)

	router.PUT("entry/update/:id", routes.UpdateEntry)
	router.PUT("ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("entry/delete/:id", routes.DeleteEntry)

	router.Run(":" + port)
}