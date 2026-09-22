package habit

import "time"

//Habit хранит описание самой привычки

type Habit struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsActive    bool      `db:"is_active"`
}

type Status string

const (
	StatusCompleted Status = "completed"
	StatusPartial   Status = "partial"
	StatusSkipped   Status = "skipped"
)

// HabitLog хранит одну конкретную отметку по дате

type HabitLog struct {
	HabitID int       `db:"habit_id"`
	Status  Status    `db:"status"` //completed, partial, skipped
	Date    time.Time `db:"date"`
	Note    *string   `db:"note"`
}
