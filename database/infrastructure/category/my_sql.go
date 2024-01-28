package infrastructure

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
)

type MySqlCategoryRepository struct {
	db *sql.DB
}

func NewMySqlCategoryRepository(db *sql.DB) CategoryRepository {
	return &MySqlCategoryRepository{
		db: db,
	}
}

// TODO: Implementar metodos.
func (mySql *MySqlCategoryRepository) Create(category *domain_task.Category) (*domain_task.Category, error) {
	return nil, nil
}

func (mySql *MySqlCategoryRepository) Delete(category *domain_task.Category) (*domain_task.Category, error) {
	return nil, nil
}

func (mySql *MySqlCategoryRepository) Get(user *domain.User) (*[]domain_task.Category, error) {
	return nil, nil
}
