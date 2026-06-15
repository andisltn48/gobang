package handler

import (
	"encoding/json"
	"gobang/internal/dto"
	"gobang/internal/service"
	"gobang/internal/utils"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload")
	}

	if req.Username == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Username is required")
	}
	if req.Email == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Email is required")
	}
	if req.Password == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Password is required")
	}

	_, err = h.authService.Register(r.Context(), req)

	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JsonResponse(w, http.StatusCreated, "User registered successfully")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Email == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	if req.Password == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Password is required")
		return
	}

	token, err := h.authService.Login(r.Context(), req)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusOK, map[string]string{"token": token})
}
 