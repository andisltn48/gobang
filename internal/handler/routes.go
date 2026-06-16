package handler

import (
	"gobang/internal/utils"
	"net/http"
)

func AuthRoutes(mux *http.ServeMux, authHandler *AuthHandler) {
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
	me := http.HandlerFunc(authHandler.Me)
	mux.Handle("GET /me", utils.JWTMiddleware(me))
}

func UserRoutes(mux *http.ServeMux, userHandler *UserHandler) {
	getUsers := http.HandlerFunc(userHandler.GetUsers)
	mux.Handle("GET /users", utils.JWTMiddleware(getUsers))
}