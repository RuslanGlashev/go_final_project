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
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		err := db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]interface{}{}, http.StatusOK)
		return
	}
	dateParse, err := time.Parse(layout, task.Date)
	if err != nil {
		writeJson(w, map[string]string{"error": "Неверная дата задачи"}, http.StatusInternalServerError)
		return
	}
	// Рассчитываем следующую дату
	next, err := NextDate(dateParse, task.Date, task.Repeat)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}
	err = db.UpdateDate(next, id)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJson(w, map[string]interface{}{}, http.StatusOK)
}
