package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/facundocarballo/microservices-todo-list/api/utils"
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
		fmt.Println(utils.ERR_EXECUTING_SP, err.Error())
		return nil, err
	}

	err = mySql.db.QueryRow("SELECT @newId, @newCreatedAt").Scan(&newId, &newCreatedAt)
	if err != nil {
		fmt.Println(utils.ERR_QUERYING, err.Error())
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
	rows, err := mySql.db.Query("SELECT username, email, photo_url, created_at FROM User WHERE id = (?)", id)
	if err != nil {
		fmt.Println(utils.ERR_QUERY, err.Error())
		return nil, err
	}
	defer rows.Close()

	var user domain.User
	user.Id = id
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Email, &user.PhotoUrl, &user.CreatedAt)
		if err != nil {
			fmt.Println(utils.ERR_END_QUERY, err.Error())
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error at the end of the query. ", err.Error())
		return nil, err
	}

	return &user, nil
}
