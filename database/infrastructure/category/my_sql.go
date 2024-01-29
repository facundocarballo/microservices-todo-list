package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/facundocarballo/microservices-todo-list/api/utils"
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

func (mySql *MySqlCategoryRepository) Create(category *domain_task.Category) (*domain_task.Category, error) {
	var newId int
	var newCreatedAt []uint8

	_, err := mySql.db.Exec(
		"CALL CreateCategory(?, ?, ?, ?, ?, @newId, @newCreatedAt)",
		category.Name, category.Description, category.PhotoUrl, category.ParentId, category.UserId,
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

	category.Id = newId
	category.CreatedAt = newCreatedAt

	return category, nil
}

func (mySql *MySqlCategoryRepository) Delete(category *domain_task.Category) (*domain_task.Category, error) {
	return nil, nil
}

func (mySql *MySqlCategoryRepository) Get(user *domain.User) ([]*domain_task.Category, error) {
	rows, err := mySql.db.Query("CALL GetAllCategoriesFromUser(?)", user.Id)
	if err != nil {
		fmt.Println(utils.ERR_QUERY, err.Error())
		return nil, err
	}

	var categories []*domain_task.Category
	for rows.Next() {
		var category domain_task.Category
		category.UserId = user.Id

		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.PhotoUrl,
			&category.CreatedAt,
			&category.ParentId,
			&category.UserId,
		)
		if err != nil {
			fmt.Println(utils.ERR_ITERATING_QUERY, err.Error())
			return nil, err
		}

		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(utils.ERR_END_QUERY, err.Error())
		return nil, err
	}

	return categories, nil
}
