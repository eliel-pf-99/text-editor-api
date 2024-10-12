package main

import (
	"log"
	"net/http"
	"os"
	"server/db"
	"server/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_uri := os.Getenv("connectiondb")
	db_name := os.Getenv("dbname")
	collection_name := os.Getenv("collectionuser")
	db, _ := db.NewConn(db_uri)
	defer db.Close()

	repository := users.NewRepository(db.GetDB(), db_name, collection_name)
	service := users.NewService(repository)
	handler := users.NewHandler(service)

	r := gin.Default()
	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	r.GET("/protect", handler.Auth, func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		ctx.JSON(http.StatusOK, gin.H{"name": user})
	})
	r.Run()

}
