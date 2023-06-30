package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	driverName = "postgres"
	sslMode    = "?sslmode=disable"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

func GetDBSession(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, dbURL+sslMode)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	return db, nil
}
