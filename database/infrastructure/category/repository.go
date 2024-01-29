package infrastructure

import (
	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
)

type CategoryRepository interface {
	Create(category *domain_task.Category) (*domain_task.Category, error)
	Delete(category *domain_task.Category) (*domain_task.Category, error)
	Get(user *domain.User) ([]*domain_task.Category, error)
}
