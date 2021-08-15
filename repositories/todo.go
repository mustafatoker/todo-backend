package repositories

import (
	"database/sql"
	"golang-todo-list/models"
	"log"
	"time"
)

func Create(db *sql.DB, name string) (int, error) {
	createdAt := time.Now()

	stmt, err := db.Prepare(`INSERT INTO todos (name, created_at) VALUES ($1,$2) RETURNING id`)

	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(name, createdAt)

	var id int
	err = row.Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id, err
}

func GetAll(db *sql.DB) ([]*models.Todo, error) {
	rows, err := db.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	result := make([]*models.Todo, 0)

	for rows.Next() {
		todo := new(models.Todo)
		if err := rows.Scan(&todo.Id, &todo.Name, &todo.CreatedAt); err != nil {
			log.Fatal(err)
		}

		result = append(result, todo)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func Delete(db *sql.DB, id string) (int64, error) {
	res, err := db.Exec(`DELETE FROM todos WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}

	numDeleted, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return numDeleted, err
}
