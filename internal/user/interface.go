package user

import "go-marketplace/internal/models"

type UserRepository interface {
	AddUser(user models.User) (models.User, error)
	FindByEmail(email string) (*models.User, error)
}
