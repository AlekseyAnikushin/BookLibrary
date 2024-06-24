package main

import (
	"log"
	"net/http"

	db "book_library/pkg/database"
	handlers "book_library/pkg/handlers"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handlers.GetMux(),
	}

	dbErr := db.InitDB()
	if dbErr != nil {
		log.Fatalln("Database connection error: ", dbErr.Error())
		return
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln("Не удалось запустить сервер: ", err.Error())
	}
}
