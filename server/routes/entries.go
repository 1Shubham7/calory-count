package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = openCollection(Client, "calories")

func AddEntry(c *gin.Context){

}


func GetEntries(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	if err  = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return 
	}

	defer cancel()
	fmt.Println(err)

}


func GetEntryById(c *gin.Context) {

}


func GetEntriesByIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.Context){

}


func UpdateIngredient(c *gin.Context){

}

func DeleteEntry(c *gin.Context){
	entryID := c.Params.ByName("id")
	docID,_ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	result,err := entryCollection.DeleteOne(ctx, bson.M{"_id":docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}