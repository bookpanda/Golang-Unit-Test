package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
    Service *UserService
}

func NewUserController(service *UserService) *UserController {
    return &UserController{Service: service}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    userId := r.URL.Query().Get("userId")
    if userId == "" {
        http.Error(w, "userId is required", http.StatusBadRequest)
        return
    }

    user, err := uc.Service.GetUser(userId)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
