package repository

import (
	todo "Go_new_api"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type todoListPostgres struct {
	db *sqlx.DB
}

func newTodoListPostgres(db *sqlx.DB) *todoListPostgres {
	return &todoListPostgres{db: db}
}

func (r *todoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSER INTO %s (title, descriptiom) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUser
}
