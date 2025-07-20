package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "Negr Akif", Email: "qaziyem@example.com"},
    {ID: 2, Name: "Hummet", Email: "sasalele@example.com"},
}

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })

    r.Run(":8080")
}
