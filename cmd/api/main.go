package main

import (
	"gobang/internal/db"
	"gobang/internal/handler"
	"gobang/internal/repository"
	"gobang/internal/service"
	"log"
	"net/http"
)

func main() {
	db := db.Connect()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	// memberRepo := repository.NewMemberRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()
	handler.AuthRoutes(mux, authHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}