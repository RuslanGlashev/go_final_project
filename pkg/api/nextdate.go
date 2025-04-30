package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const layout = "20060102"

// afterNow проверяет, является ли date больше now
func afterNow(date, now time.Time) bool {
	return date.After(now)
}

// Основная функция, которая вычисляет следующую дату
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Парсим дату в формате "20060102"
	date, err := time.Parse(layout, dstart)
	if err != nil {
		return "", fmt.Errorf("Некорректный формат даты: %v", err)
	}

	// Разбиваем повторение на составляющие
	rules := strings.Split(repeat, " ")

	if len(rules) == 0 {
		return "", fmt.Errorf("Правило не указано, задача будет удалена")
	}

	switch rules[0] {
	case "d":
		// Если правило касается дней
		if len(rules) < 2 {
			return "", fmt.Errorf("Не указан интервал 'd'")
		}
		interval, err := strconv.Atoi(rules[1])
		if err != nil || interval <= 0 || interval > 400 {
			return "", fmt.Errorf("Недопустимое значение интервала: %v", err)
		}

		for {
			date = date.AddDate(0, 0, interval) // Добавляем дни
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(layout), nil

	case "y":
		// Если правило ежегодое

		date = date.AddDate(1, 0, 0) // Добавляем 1 год
		if afterNow(date, now) {
			break
		}
		return date.Format(layout), nil

	case "m", "w":
		return "", fmt.Errorf("Некорректный запрос")

	default:
		return "", fmt.Errorf("Неизвестное правило")

	}

	// Преобразуем дату в строку в формате 20060102 и выводим результат
	return date.Format(layout), nil
}

// nextDayHandler обрабатывает запросы к /api/nextdate
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	now := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	nowParse, _ := time.Parse(layout, now)

	nextDate, err := NextDate(nowParse, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(nextDate))
}
