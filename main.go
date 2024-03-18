package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Placeholder endpoint
    r.GET("/api/v1/placeholder", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "This is a placeholder endpoint"})
    })

    // Start the server
    r.Run(":3000")
}