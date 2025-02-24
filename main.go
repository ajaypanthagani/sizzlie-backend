package main

import (
	"database/sql"
	"log"

	handler "github.com/ajaypanthagani/sizzlie/handler/user"
	repository "github.com/ajaypanthagani/sizzlie/internal/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:MsendPoint@5140@tcp(127.0.0.1:3306)/sizzlie"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close() // Initialize the repository
	userRepo := repository.NewUserRepository(db)

	// Initialize the handler
	userHandler := handler.NewUserHandler(userRepo)

	router := gin.Default()

	router.POST("/signup", func(c *gin.Context) {
		userHandler.Create(c)
	})

	router.Run(":8000")

}
