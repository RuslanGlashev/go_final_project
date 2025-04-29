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
		task.Date = now.Format("20060102")
	}

	// Проверяем, корректна ли указанная дата
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("Дата представлена в неправильном формате: %v", err)
	}

	// если сегодня (now) больше task.Date (t)
	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format("20060102")
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("Не удалось вычислить следующую дату по правилу повторения: %v", err)
			}
			task.Date = next
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, map[string]string{"error": "Ошибка десериализации JSON"})
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверка даты задачи
	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	writeJson(w, map[string]string{"id": fmt.Sprintf("%d", id)})
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}
