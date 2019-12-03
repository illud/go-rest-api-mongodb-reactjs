package routes

import (
	"context"
	"fmt"
	"log"

	db "../db"
	models "../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewTask(c *gin.Context) {
	var taskBody models.Task
	c.BindJSON(&taskBody)
	insert := models.Task{taskBody.TITLE, taskBody.DESCRIPTION, taskBody.DATE}
	collection := db.CNX.Database("tasks").Collection("task")
	insertResult, err := collection.InsertOne(context.TODO(), insert)
	if err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{
			"message": "Error",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Inserted",
		})
		fmt.Println(insertResult)
	}
}

func GetTasks(c *gin.Context) {

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(100)
	collection := db.CNX.Database("tasks").Collection("task")
	// Here's an array in which you can store the decoded documents
	var results []*models.TaskGet

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into whch the single document can be decoded
		var elem models.TaskGet
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	c.JSON(200, gin.H{
		"tasks": results,
	})
}

func DeleteTask(c *gin.Context) {
	//https://kb.objectrocket.com/mongo-db/how-to-delete-mongodb-documents-using-the-golang-driver-443
	var taskBody models.TaskGet
	c.BindJSON(&taskBody)

	collection := db.CNX.Database("tasks").Collection("task")
	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex(taskBody.ID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	c.JSON(200, gin.H{
		"tasks": "Deleted",
	})
}
