package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	go server.NotifyManager.CheckEventsAndNotify(ctx)
	defer cancel()
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	sig := <-sigs
	// 	cancel()
	// 	log.Fatal(sig)
	// }()

	router.POST("/register", server.CreateUserHandler)

	auth := router.Group("/")

	auth.Use(middleware.Auntification())

	auth.GET("/profile", server.GetUserHandler)

	auth.POST("/event", server.AddEventHandler)

	auth.POST("/sub", server.AddSubscriptionHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
