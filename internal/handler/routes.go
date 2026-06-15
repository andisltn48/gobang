package handler

import (
	"net/http"
)

func AuthRoutes(mux *http.ServeMux, authHandler *AuthHandler) {
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
}