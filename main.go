package main

import (
	"database/sql"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
)

func dbClose(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Ошибка при закрытии базы данных: %v", err)
	}
}

func main() {
	database, err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}
	defer dbClose(database)
	server.Run()
}
