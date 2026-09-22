package habit

import (
	"errors"
)

func (h Habit) Validate() error {

	if h.Title == "" {
		return errors.New("Название не может быть пустым")
	}

	return nil
}

func IsValidStatus(status Status) bool {
	return status == StatusCompleted || status == StatusPartial || status == StatusSkipped
}

func (h HabitLog) Validate() error {

	if !IsValidStatus(h.Status) {
		return errors.New("Некорректный статус")
	}

	if h.Date.IsZero() {
		return errors.New("Дата не может быть пустой")
	}

	return nil
}
