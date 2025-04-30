# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

# Проекта
Проект веб-браузер. Планировщик задач который добавляет, изменяет, удаляет, отмечает выполненные задачи, вычисляет следующую дату для задач.

# Список заданий со звёздочкой:
- Создание Docker-образа.

# Версия Go 1.23.5

# Откройте страницу браузера по ссылке: `http://localhost:7540`

# .env :
TODO_PORT=7540
TODO_DBFILE="../scheduler.db"
TODO_PASSWORD=1234

# В tests/settings.go следует использовать:
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = false
var Search = false
var Token = ``


# Docker:

1. docker build -t go_final_project .

2. docker run -d -p 7540:7540 -v /data/scheduler.db go_final_project