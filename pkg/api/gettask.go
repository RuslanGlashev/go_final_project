package api

import (
	"net/http"

	"go1f/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем идентификатор задачи из параметров запроса
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	writeJson(w, task)
}
