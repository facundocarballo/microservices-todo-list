package infrastructure

import (
	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
)

type TaskRepository interface {
	Create(task *domain_task.Task) (*domain_task.Task, error)
	Update(task *domain_task.Task) (*domain_task.Task, error)
	Complete(task *domain_task.Task) (*domain_task.TaskCompleted, error)
	Delete(task *domain_task.Task) (*domain_task.TaskDeleted, error)
	Get(user *domain.User) ([]*domain_task.Task, error)
}
