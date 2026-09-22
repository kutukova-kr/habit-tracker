package habit

import (
	"database/sql"
	"time"
)

type habitLogRepository struct {
	db *sql.DB
}

//func NewHabitLogRepository(db *sql.DB) HabitLogRepository {
//	return &habitLogRepository{db: db}
//}

func (r *habitLogRepository) CreateHabitLog(log *HabitLog) error {
	_, err := r.db.Exec(
		`INSERT INTO habit_logs (habit_id, status, date, note)
         VALUES ($1, $2, $3, $4)`,
		log.HabitID,
		log.Status,
		log.Date,
		log.Note,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *habitLogRepository) GetHabitLogByHabitAndDate(habitID int, date time.Time) (*HabitLog, error) {
	var log HabitLog

	err := r.db.QueryRow(
		`SELECT habit_id, status, date, note
		FROM habit_logs
		WHERE habit_id = $1 AND date = $2`,
		habitID,
		date,
	).Scan(
		&log.HabitID,
		&log.Status,
		&log.Date,
		&log.Note,
	)
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (r *habitLogRepository) ListHabitLogsByHabit(habitID int) ([]HabitLog, error) {
	var logs []HabitLog

	rows, err := r.db.Query(
		`SELECT habit_id, status, date, note
		FROM habit_logs
		WHERE habit_id = $1`,
		habitID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log HabitLog

		err := rows.Scan(
			&log.HabitID,
			&log.Status,
			&log.Date,
			&log.Note,
		)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *habitLogRepository) DeleteHabitLog(habitID int, date time.Time) error {
	_, err := r.db.Exec(
		`DELETE FROM habit_logs
		WHERE habit_id = $1 AND date = $2`,
		habitID,
		date,
	)
	return err
}
