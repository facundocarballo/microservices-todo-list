package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
)

type MySqlTaskRepository struct {
	db *sql.DB
}

func NewMySqlTaskRepository(db *sql.DB) TaskRepository {
	return &MySqlTaskRepository{
		db: db,
	}
}

func (mySql *MySqlTaskRepository) Create(task *domain_task.Task) (*domain_task.Task, error) {
	var newId int
	var newCreatedAt []uint8

	_, err := mySql.db.Exec(
		"CALL CreateTask(?, ?, ?, @newId, @newCreatedAt)",
		task.Name, task.Description, task.CategoryId,
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

	task.Id = newId
	task.CreatedAt = newCreatedAt

	return nil, nil
}

func (mySql *MySqlTaskRepository) Update(task *domain_task.Task) (*domain_task.Task, error) {
	return nil, nil
}

func (mySql *MySqlTaskRepository) Complete(task *domain_task.Task) (*domain_task.TaskCompleted, error) {
	var newId int
	var newCompletedAt []uint8

	_, err := mySql.db.Exec(
		"CALL CompleteTask(?, @newId, @newCompletedAt)",
		task.Id,
	)
	if err != nil {
		fmt.Println("Error: executing the stored procedure.\n", err.Error())
		return nil, err
	}

	err = mySql.db.QueryRow("SELECT @newId, @newCreatedAt").Scan(&newId, &newCompletedAt)
	if err != nil {
		fmt.Println("Error: executing query row.\n", err.Error())
		return nil, err
	}

	return &domain_task.TaskCompleted{
		Id:          newId,
		TaskId:      task.Id,
		CompletedAt: newCompletedAt,
	}, nil
}

func (mySql *MySqlTaskRepository) Delete(task *domain_task.Task) (*domain_task.TaskDeleted, error) {
	var newId int
	var newDeletedAt []uint8

	_, err := mySql.db.Exec(
		"CALL DeleteTask(?, @newId, @newDeletedAt)",
		task.Id,
	)
	if err != nil {
		fmt.Println("Error: executing the stored procedure.\n", err.Error())
		return nil, err
	}

	err = mySql.db.QueryRow("SELECT @newId, @newDeletedAt").Scan(&newId, &newDeletedAt)
	if err != nil {
		fmt.Println("Error: executing query row.\n", err.Error())
		return nil, err
	}

	return &domain_task.TaskDeleted{
		Id:        newId,
		TaskId:    task.Id,
		DeletedAt: newDeletedAt,
	}, nil
}

func (mySql *MySqlTaskRepository) Get(user *domain.User) ([]*domain_task.Task, error) {
	rows, err := mySql.db.Query("CALL GetTaskToDoFromUser(?)", user.Id)
	if err != nil {
		fmt.Println("Error querying. ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain_task.Task

	for rows.Next() {
		var task domain_task.Task
		task.UserId = user.Id

		err := rows.Scan(&task.Name, &task.Description, &task.CreatedAt, &task.CategoryId)
		if err != nil {
			fmt.Println("Error iterating the query. ", err.Error())
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error at the end of the query. ", err.Error())
		return nil, err
	}

	return tasks, nil
}
