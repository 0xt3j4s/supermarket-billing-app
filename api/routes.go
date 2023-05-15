package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Define routes
    r.GET("/bills", getAllBills)
    r.POST("/bills", createBill)
    r.GET("/bills/:id", getBill)

    return r
}
