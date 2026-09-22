package habit

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type HabitHandler struct {
	svc *HabitService
}

func NewHabitHandler(svc *HabitService) http.Handler {
	return &HabitHandler{svc: svc}
}

func (h *HabitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/habits":
		switch r.Method {
		case http.MethodGet:
			h.listHabits(w, r)
		case http.MethodPost:
			h.createHabit(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	default:
		// маршруты вида /habits/{id}
		if len(r.URL.Path) > 8 && r.URL.Path[:8] == "/habits/" {
			idStr := r.URL.Path[8:]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid habit id", http.StatusBadRequest)
				return
			}

			switch r.Method {
			case http.MethodGet:
				h.getHabitByID(w, r, id)
			case http.MethodPut, http.MethodPatch:
				h.updateHabit(w, r, id)
			case http.MethodDelete:
				h.deleteHabit(w, r, id)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (h *HabitHandler) listHabits(w http.ResponseWriter, r *http.Request) {
	habits, err := h.svc.ListHabits()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(habits)
}

func (h *HabitHandler) createHabit(w http.ResponseWriter, r *http.Request) {
	var input Habit
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.svc.CreateHabit(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *HabitHandler) getHabitByID(w http.ResponseWriter, r *http.Request, id int) {
	habit, err := h.svc.GetHabitByID(id)
	if err != nil {
		if err == ErrInvalidHabitID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(habit)
}

func (h *HabitHandler) updateHabit(w http.ResponseWriter, r *http.Request, id int) {
	var input Habit
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updated, err := h.svc.UpdateHabit(id, input)
	if err != nil {
		if err == ErrInvalidHabitID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *HabitHandler) deleteHabit(w http.ResponseWriter, r *http.Request, id int) {
	err := h.svc.DeleteHabit(id)
	if err != nil {
		if err == ErrInvalidHabitID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
