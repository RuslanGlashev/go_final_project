// pkg/api/addtask.go
package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/db"
	"net/http"
	"time"
)

// checkDate проверяет корректность даты задачи и обновляет её при необходимости
func checkDate(task *db.Task) error {
	now := time.Now()

	// Если дата не указана, присваиваем текущую дату
	if task.Date == "" {
		task.Date = now.Format(layout)
	}

	tNow, _ := time.Parse(layout, now.Format(layout))
	// Проверяем, корректна ли указанная дата
	t, err := time.Parse(layout, task.Date)
	if err != nil {
		return fmt.Errorf("Дата представлена в неправильном формате: %v", err)
	}

	// если сегодня (now) больше task.Date (t)
	if afterNow(tNow, t) {
		if task.Repeat == "" {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format(layout)
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("Не удалось вычислить следующую дату по правилу повторения: %v", err)
			}
			task.Date = next
		}
	}
	test, err := time.Parse(layout, task.Date)
	if err != nil {
		return fmt.Errorf("Дата представлена в неправильном формате: %v", err)
	}
	if afterNow(tNow, test) {
		next, err := NextDate(now, task.Date, task.Repeat)
		task.Date = next
		return err
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, map[string]string{"error": "Ошибка десериализации JSON"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	// Проверка даты задачи
	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("Ошибка добавления задачи в базу данных: %v", err)}, http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]int64{"id": id}, http.StatusOK)
}

func writeJson(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка при обработке JSON", http.StatusInternalServerError)
		return
	}
}
