package main

import (
	"log"
	"net/http"

	"book_library/internal/handlers"
	db "book_library/internal/storages"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handlers.GetMux(),
	}

	dbErr := db.InitDB()
	if dbErr != nil {
		log.Fatalln("Database connection error: ", dbErr.Error())
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln("Не удалось запустить сервер: ", err.Error())
	}
}
