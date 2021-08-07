package main

import (
	"database/sql"
	"demo"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// var DB *sql.DB
type Resource struct {
	db *sql.DB
}

func main() {
	// DB = createDatabaseConnection()
	r := demo.CreateDatabaseConnection()
	// users, _ := getAllUsers(db)
	// fmt.Printf("%+v", users)

	router := gin.New()
	route := router.Group("/api/v1")
	route.GET("/user", handleGetUsers(r))
	router.Run(":8080")

}

// Closure
func handleGetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := demo.GetAllUsers(db)
		c.JSON(http.StatusOK, users)
	}
}
