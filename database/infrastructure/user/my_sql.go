package infrastructure

import (
	"database/sql"
	"fmt"

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

func (mySql *MySqlUserRepository) Create(user *domain.User) (*domain.User, error) {
	var newId int
	var newCreatedAt []uint8
	_, err := mySql.db.Exec(
		"CALL CreateUser(?, ?, ?, ?, @newId, @newCreatedAt)",
		user.Username, user.Email, user.Password, user.PhotoUrl,
	)
	if err != nil {
		fmt.Println("Error: executing the stored procedure.\n", err.Error())
		return nil, err
	}

	err = mySql.db.QueryRow("SELECT @newId, @newCreatedAt").Scan(&newId, &newCreatedAt)
	if err != nil {
		fmt.Println("Error: executing query row.\n", err.Error())
		return nil, err
	}

	user.Id = newId
	user.CreatedAt = newCreatedAt
	return user, nil
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
