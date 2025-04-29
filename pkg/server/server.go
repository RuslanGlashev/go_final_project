package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go1f/pkg/api"

	"github.com/joho/godotenv"
)

func Run() {
	// инициализация API-обработчиков
	api.Init()

	webDir := "./web/"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	key := "TODO_PORT"
	port := os.Getenv(key)
	log.Println("Сервер запускается. Port:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
