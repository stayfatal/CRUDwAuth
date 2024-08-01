package main

import (
	"log"
	"net/http"
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server, err := handlers.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/register", server.CreateUserHandler)

	auth := router.Group("/")

	auth.Use(middleware.Auntification())

	auth.GET("/profile", server.GetUserHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
