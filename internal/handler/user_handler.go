package handler

import (
	"gobang/internal/dto"
	"gobang/internal/service"
	"gobang/internal/utils"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers(r.Context())
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	userResponse := make([]dto.UserResponseDTO, 0)
	for _, user := range users {
		userData := dto.UserResponseDTO{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		// Only populate member if exists
		if user.Member != nil {
			userData.Member = &dto.MemberResponseDTO{
				ID:        user.Member.ID,
				FirstName: user.Member.FirstName,
				LastName:  user.Member.LastName,
				Saldo:     user.Member.Saldo,
			}
		}

		userResponse = append(userResponse, userData)
	}

	result := map[string]interface{}{
		"data": userResponse,
	}

	utils.JsonResponse(w, http.StatusOK, result)
}
