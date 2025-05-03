package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go1f/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, map[string]string{"error": "Некорректные данные"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	// Проверка даты задачи
	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("%v", err)}, http.StatusBadRequest)
		return
	}

	// Обновляем задачу в базе данных
	if err := db.UpdateTask(&task); err != nil {
		writeJson(w, map[string]string{"error": "Задача не найдена"}, http.StatusNotFound)
		return
	}

	// Возвращаем пустой JSON
	writeJson(w, map[string]interface{}{}, http.StatusOK)
}
