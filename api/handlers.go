package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/0xt3j4s/supermarket-billing-app/data"
	"github.com/0xt3j4s/supermarket-billing-app/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllBills(c *gin.Context) {

	// Define an empty slice to store the retrieved bills
	var bills []data.Bill

	// Retrieve all bills from the database
	billsCollection := db.GetBillsCollection()
	cursor, err := billsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the retrieved bills and append them to the slice
	for cursor.Next(context.Background()) {
		var bill data.Bill
		err := cursor.Decode(&bill)
		if err != nil {
			log.Fatal(err)
		}
		bills = append(bills, bill)
	}

	// Convert the bills slice to JSON with indentation
    jsonBytes, err := json.MarshalIndent(bills, "", "  ")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }
    
    // Set the response header and write the JSON to the response
    c.Header("Content-Type", "application/json")
    c.Writer.WriteHeader(http.StatusOK)
    c.Writer.Write(jsonBytes)
}

func createBill(c *gin.Context) {

    var newBill data.Bill
    if err := c.ShouldBindJSON(&newBill); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Create a slice to store the items
    var items []data.Item

    // Loop through the request payload to get the items
    for _, item := range newBill.Items {
        items = append(items, item)
    }

    // Add the items to the new bill
    newBill.Items = items

    // Insert the new bill into the database
    billsCollection := db.GetBillsCollection()
    _, err := billsCollection.InsertOne(context.Background(), newBill)
    if err != nil {
        log.Fatal(err)
    }

    // Convert the newBill to JSON with indentation
    jsonBytes, err := json.MarshalIndent(newBill, "", "  ")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Set the response header and write the JSON to the response
    c.Header("Content-Type", "application/json")
    c.Writer.WriteHeader(http.StatusCreated)
    c.Writer.Write(jsonBytes)
}

func getBill(c *gin.Context) {

	billIDStr := c.Param("id")

	// Convert the billID to an ObjectId
    billID, err := strconv.Atoi(billIDStr)
    if err != nil {
		log.Println("Error converting bill ID to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	log.Print("billid: ", billID)

	billsCollection := db.GetBillsCollection()
	
	var bill data.Bill
	err = billsCollection.FindOne(context.Background(), bson.M{"id": billID}).Decode(&bill)
	if err != nil {
		log.Fatal(err)
	}
	
	// Convert the bill to JSON with indentation
    jsonBytes, err := json.MarshalIndent(bill, "", "  ")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Set the response header and write the JSON to the response
    c.Header("Content-Type", "application/json")
    c.Writer.WriteHeader(http.StatusOK)
    c.Writer.Write(jsonBytes)
}


func updateBill(c *gin.Context) {

    billIDStr := c.Param("id")

	// Convert the billID to an ObjectId
    billID, err := strconv.Atoi(billIDStr)
    if err != nil {
		log.Println("Error converting bill ID to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}


	// Retrieve the existing bill from the database
    billsCollection := db.GetBillsCollection()
    filter := bson.M{"id": billID}
    var existingBill data.Bill
    err = billsCollection.FindOne(context.Background(), filter).Decode(&existingBill)
    if err != nil {
        log.Fatal(err)
    }

    var updatedBill data.Bill
    if err := c.ShouldBindJSON(&updatedBill); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

	// Exclude the ID field from the update
    updatedBill.ID = existingBill.ID


    update := bson.M{"$set": updatedBill}
    _, err = billsCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

    c.JSON(http.StatusOK, updatedBill)
}

func deleteBill(c *gin.Context) {

	billIDStr := c.Param("id")

	// Convert the billID to an ObjectId
    billID, err := strconv.Atoi(billIDStr)
    if err != nil {
		log.Println("Error converting bill ID to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

    // Delete the bill from the database
    billsCollection := db.GetBillsCollection()
    filter := bson.M{"id": billID}
    _, err = billsCollection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Fatal(err)
    }

    c.JSON(http.StatusOK, gin.H{"message": "Bill deleted"})
}