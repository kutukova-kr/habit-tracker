package habit

import "time"

type HabitRepository interface {
	CreateHabit(habit *Habit) (*Habit, error)
	GetHabitByID(id int) (*Habit, error)
	UpdateHabit(id int, habit Habit) (*Habit, error)
	DeleteHabit(id int) error
	ListHabits() ([]Habit, error)
}

type HabitLogRepository interface {
	CreateHabitLog(log *HabitLog) error
	GetHabitLogByHabitAndDate(habitID int, date time.Time) (*HabitLog, error)
	ListHabitLogsByHabit(habitID int) ([]HabitLog, error)
	DeleteHabitLog(habitID int, date time.Time) error
}
