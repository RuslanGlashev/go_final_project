package api

import (
	"go1f/pkg/db"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Repeat == "" {
		err := db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error": err.Error()})
			return
		}
		writeJson(w, map[string]interface{}{})
		return
	}
	dateParse, err := time.Parse("20060102", task.Date)
	if err != nil {
		writeJson(w, map[string]string{"error": "Неверная дата задачи"})
		return
	}
	// Рассчитываем следующую дату
	next, err := NextDate(dateParse, task.Date, task.Repeat)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	err = db.UpdateDate(next, id)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, map[string]interface{}{})
}
