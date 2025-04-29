package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL,
    comment TEXT,
    repeat VARCHAR(128)
);
CREATE INDEX idx_date ON scheduler(date);
`

var db *sql.DB

func Init(dbFile string) error {
	// Проверка существования файла базы данных
	_, err := os.Stat(dbFile)
	install := false
	if err != nil {
		install = true
	}

	// Открываем базу данных
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Если файл не существовал, создаем таблицы
	if install {
		_, err := db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}
