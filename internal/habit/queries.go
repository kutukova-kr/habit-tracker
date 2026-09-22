package habit

import "database/sql"

func CreateHabit(tx *sql.Tx, h *Habit) error {
	_, err := tx.Exec(
		`INSERT INTO habits (title, description, created_at, updated_at, is_active)
         VALUES ($1, $2, $3, $4, $5)`,
		h.Title,
		h.Description,
		h.CreatedAt,
		h.UpdatedAt,
		h.IsActive,
	)
	return err
}

func CreateHabitLog(tx *sql.Tx, log *HabitLog) error {
	_, err := tx.Exec(
		`INSERT INTO habit_logs (habit_id, status, date, note)
         VALUES ($1, $2, $3, $4)`,
		log.HabitID,
		log.Status,
		log.Date,
		log.Note,
	)
	return err
}
