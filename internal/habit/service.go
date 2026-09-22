package habit

import (
	"errors"
	"time"
)

type HabitService struct {
	habitRepo    HabitRepository
	habitLogRepo HabitLogRepository
}

func NewHabitService(habitRepo HabitRepository, habitLogRepo HabitLogRepository) *HabitService {
	return &HabitService{
		habitRepo:    habitRepo,
		habitLogRepo: habitLogRepo,
	}
}

func (s *HabitService) CreateHabit(habit Habit) (*Habit, error) {
	if err := habit.Validate(); err != nil {
		return nil, err
	}

	createdHabit, err := s.habitRepo.CreateHabit(&habit)
	if err != nil {
		return nil, err
	}

	return createdHabit, nil
}

func (s *HabitService) ListHabits() ([]Habit, error) {
	return s.habitRepo.ListHabits()
}

func (s *HabitService) CreateHabitLog(log *HabitLog) error {
	if err := log.Validate(); err != nil {
		return err
	}

	return s.habitLogRepo.CreateHabitLog(log)
}

var ErrInvalidHabitID = errors.New("invalid habit id")

func (s *HabitService) GetHabitByID(id int) (*Habit, error) {
	if id <= 0 {
		return nil, ErrInvalidHabitID
	}

	return s.habitRepo.GetHabitByID(id)
}

func (s *HabitService) UpdateHabit(id int, input Habit) (*Habit, error) {
	if id <= 0 {
		return nil, ErrInvalidHabitID
	}

	habit, err := s.habitRepo.GetHabitByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != "" {
		habit.Title = input.Title
	}

	if input.Description != "" {
		habit.Description = input.Description
	}

	if err := habit.Validate(); err != nil {
		return nil, err
	}

	return s.habitRepo.UpdateHabit(id, *habit)
}

func (s *HabitService) DeleteHabit(id int) error {
	if id <= 0 {
		return ErrInvalidHabitID
	}

	_, err := s.habitRepo.GetHabitByID(id)
	if err != nil {
		return err
	}

	return s.habitRepo.DeleteHabit(id)
}

func (s *HabitService) GetHabitLogByHabitAndDate(habitID int, date time.Time) (*HabitLog, error) {
	if habitID <= 0 {
		return nil, ErrInvalidHabitID
	}

	return s.habitLogRepo.GetHabitLogByHabitAndDate(habitID, date)
}

func (s *HabitService) ListHabitLogsByHabit(habitID int) ([]HabitLog, error) {
	if habitID <= 0 {
		return nil, ErrInvalidHabitID
	}

	return s.habitLogRepo.ListHabitLogsByHabit(habitID)
}

func (s *HabitService) DeleteHabitLog(habitID int, date time.Time) error {
	if habitID <= 0 {
		return ErrInvalidHabitID
	}

	return s.habitLogRepo.DeleteHabitLog(habitID, date)
}
