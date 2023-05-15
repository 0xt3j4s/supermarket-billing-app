package main

import (
	"log"
	"github.com/0xt3j4s/supermarket-billing-app/api"
	"github.com/0xt3j4s/supermarket-billing-app/db"
)

func main() {
	// Connect to MongoDB
	err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	// Set up Gin router
	router := api.SetupRouter()

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
