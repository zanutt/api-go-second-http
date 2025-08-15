package user

import (
	"encoding/json"
	"go-marketplace/internal/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo UserRepository
}

func NewHandler(repo UserRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = string(hashedPassword)

	createdUser, err := h.repo.AddUser(newUser)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	createdUser.Password = "" // Do not expose the password hash
	json.NewEncoder(w).Encode(createdUser)
}
