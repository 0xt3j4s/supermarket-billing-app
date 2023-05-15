package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/0xt3j4s/supermarket-billing-app/db"
	"github.com/0xt3j4s/supermarket-billing-app/constants"
)

func getAllBills(c *gin.Context) {
    // Implement logic to retrieve all bills from the database
    // and return them as a response

	// Define an empty slice to store the retrieved bills
	var bills []Bill

	// Retrieve all bills from the database
	billsCollection := db.GetBillsCollection()
	cursor, err := billsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the retrieved bills and append them to the slice
	for cursor.Next(context.Background()) {
		var bill Bill
		err := cursor.Decode(&bill)
		if err != nil {
			log.Fatal(err)
		}
		bills = append(bills, bill)
	}

	c.JSON(http.StatusOK, bills)

    // c.JSON(http.StatusOK, gin.H{
    //     "message": "Get All Bills",
    // })
}

func createBill(c *gin.Context) {
    // Implement logic to create a new bill based on the request payload
    // and return the created bill as a response

	var newBill Bill
	if err := c.ShouldBindJSON(&newBill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Insert the new bill into the database
	billsCollection := db.GetBillsCollection()
	_, err := billsCollection.InsertOne(context.Background(), newBill)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newBill)

    // c.JSON(http.StatusOK, gin.H{
    //     "message": "Create Bill",
    // })
}

func getBill(c *gin.Context) {
    // Implement logic to retrieve a specific bill from the database
    // based on the provided ID and return it as a response
    billID := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "Get Bill",
        "id":      billID,
    })
}
