package habit

import (
	"database/sql"
)

type habitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) HabitRepository {
	return &habitRepository{db: db}
}

func (r *habitRepository) CreateHabit(habit *Habit) (*Habit, error) {
	err := r.db.QueryRow(
		`INSERT INTO habits (title, description, is_active)
         VALUES ($1, $2, $3)
         RETURNING id, created_at, updated_at`,
		habit.Title,
		habit.Description,
		habit.IsActive,
	).Scan(&habit.ID, &habit.CreatedAt, &habit.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return habit, nil
}

func (r *habitRepository) GetHabitByID(id int) (*Habit, error) {
	var h Habit

	err := r.db.QueryRow(
		`SELECT id, title, description, created_at, updated_at, is_active
         FROM habits
         WHERE id = $1`,
		id,
	).Scan(
		&h.ID,
		&h.Title,
		&h.Description,
		&h.CreatedAt,
		&h.UpdatedAt,
		&h.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return &h, nil
}

func (r *habitRepository) UpdateHabit(id int, habit Habit) (*Habit, error) {
	var h Habit

	err := r.db.QueryRow(
		`UPDATE habits
         SET title = $1,
             description = $2,
             is_active = $3,
             updated_at = NOW()
         WHERE id = $4
         RETURNING id, title, description, created_at, updated_at, is_active`,
		habit.Title,
		habit.Description,
		habit.IsActive,
		id,
	).Scan(
		&h.ID,
		&h.Title,
		&h.Description,
		&h.CreatedAt,
		&h.UpdatedAt,
		&h.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *habitRepository) DeleteHabit(id int) error {
	_, err := r.db.Exec(
		`DELETE FROM habits WHERE id = $1`,
		id,
	)
	return err
}

func (r *habitRepository) ListHabits() ([]Habit, error) {
	rows, err := r.db.Query(
		`SELECT id, title, description, created_at, updated_at, is_active
         FROM habits`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []Habit

	for rows.Next() {
		var h Habit
		err := rows.Scan(
			&h.ID,
			&h.Title,
			&h.Description,
			&h.CreatedAt,
			&h.UpdatedAt,
			&h.IsActive,
		)
		if err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return habits, nil
}
