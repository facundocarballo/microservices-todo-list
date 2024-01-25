package infrastructure

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/domain"
)

type MySqlUserRepository struct {
	db *sql.DB
}

func NewMySqlUserRepository(db *sql.DB) UserRepository {
	return &MySqlUserRepository{
		db: db,
	}
}

func (mySql *MySqlUserRepository) Create(user *domain.User) error {

	return nil
}

func (mySql *MySqlUserRepository) Update(user *domain.User) error {
	return nil
}

func (mySql *MySqlUserRepository) Delete(user *domain.User) error {
	return nil
}

func (mySql *MySqlUserRepository) Find(id int) (*domain.User, error) {
	return nil, nil
}
