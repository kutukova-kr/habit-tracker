package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// DSN должен совпадать с POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB в docker-compose.yml
	dsn := "postgres://postgres:secret@127.0.0.1:5432/habits?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("sql.Open error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}
	fmt.Println("Connected to DB!")

	// Пример INSERT
	var id int
	err = db.QueryRow(
		`INSERT INTO habits (name, description) VALUES ($1, $2) RETURNING id`,
		"Пить воду",
		"1.5 литра в день",
	).Scan(&id)
	if err != nil {
		log.Fatal("INSERT error:", err)
	}
	fmt.Println("Inserted habit with id:", id)

	// Пример SELECT
	rows, err := db.Query(`SELECT id, name, description FROM habits`)
	if err != nil {
		log.Fatal("SELECT error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var hID int
		var name, desc sql.NullString
		if err := rows.Scan(&hID, &name, &desc); err != nil {
			log.Fatal("SCAN error:", err)
		}
		fmt.Printf("Habit: id=%d, name=%s, description=%s\n",
			hID, name.String, desc.String)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("rows.Err error:", err)
	}
}
