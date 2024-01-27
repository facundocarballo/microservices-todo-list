package infrastructure

import (
	"github.com/facundocarballo/microservices-todo-list/domain"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
	Find(id int) (*domain.User, error)
}
