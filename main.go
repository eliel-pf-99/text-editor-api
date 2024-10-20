package main

import (
	"net/http"
	"os"
	"server/db"
	"server/internal/notes"
	"server/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load("./.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db_uri := os.Getenv("connectiondb")
	db_name := os.Getenv("dbname")
	collection_user := os.Getenv("collectionuser")
	collection_note := os.Getenv("collectionnote")
	db, _ := db.NewConn(db_uri)
	defer db.Close()

	user_repository := users.NewRepository(db.GetDB(), db_name, collection_user)
	user_service := users.NewService(user_repository)
	user_handler := users.NewHandler(user_service)

	note_repository, _ := notes.NewRepository(db.GetDB(), db_name, collection_note)
	note_service := notes.NewService(note_repository)
	note_handler := notes.NewHandler(note_service)

	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"index": "Text Editor",
		})
	})

	r.POST("/signup", user_handler.Signup)
	r.POST("/login", user_handler.Login)

	r.GET("/get-notes", user_handler.Auth, note_handler.GetNotes)
	r.POST("/create-note", user_handler.Auth, note_handler.InsertNote)
	r.POST("/update-note", user_handler.Auth, note_handler.UpdateNote)
	r.POST("/get-note", user_handler.Auth, note_handler.FindNoteById)
	r.POST("/delete-note", user_handler.Auth, note_handler.DeleteNote)

	r.GET("/protect", user_handler.Auth, func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		ctx.JSON(http.StatusOK, gin.H{"name": user})
	})
	r.Run()

}
