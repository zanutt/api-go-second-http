package user

import (
	"errors"
	"go-marketplace/internal/models"
)

type InMemoryRepository struct {
	users  map[uint]models.User
	nextID uint
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users:  make(map[uint]models.User),
		nextID: 1,
	}
}

func (r *InMemoryRepository) AddUser(user models.User) (models.User, error) {
	for _, u := range r.users {
		if u.Email == user.Email {
			return models.User{}, errors.New("user already exists")
		}
	}

	user.ID = r.nextID

	r.users[r.nextID] = user

	r.nextID++

	return user, nil
}

func (r *InMemoryRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
