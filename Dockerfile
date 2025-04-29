# Используем официальный образ Golang
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o go_final_project .

# Используем образ для запуска приложения
FROM ubuntu

# Устанавливаем зависимости 
RUN apt-get update && apt-get install -y ca-certificates

# Копируем скомпилированный исполняемый файл и веб-файлы в контейнер
COPY --from=builder ./go_final_project ./go_final_project
COPY web ./web

# Устанавливаем переменные окружения
ENV TODO_PORT=7540
ENV TODO_DBFILE=./scheduler.db
ENV TODO_PASSWORD=1234

# Открываем порт
EXPOSE 7540

# Запускаем приложение
CMD ["go_final_project"]