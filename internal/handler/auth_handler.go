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
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Username == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username is required")
		return
	}
	if req.Email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	if req.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Password is required")
		return
	}

	_, err = h.authService.Register(r.Context(), req)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JsonResponse(w, http.StatusCreated, "User registered successfully")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	if req.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Password is required")
		return
	}

	token, err := h.authService.Login(r.Context(), req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(utils.UserIDKey).(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	response := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}
	utils.JsonResponse(w, http.StatusOK, response)
}

