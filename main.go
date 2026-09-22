package main

import (
	"database/sql"
	"fmt"
	"habit-tracker/internal/habit"
	"log"
	"net/http"

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

	habitRepo := habit.NewHabitRepository(db)
	habitLogRepo := habit.NewHabitLogRepository(db)

	svc := habit.NewHabitService(habitRepo, habitLogRepo)

	mux := http.NewServeMux()
	mux.Handle("/habits", habit.NewHabitHandler(svc))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
