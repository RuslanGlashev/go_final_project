package api

import (
	"net/http"
)

// Init регистрирует API-обработчики
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
	http.HandleFunc("/api/task/delete", deleteTaskHandler)

}
