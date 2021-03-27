package main

import (
	"log"
	todo_app "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("err occured while running http server: %s", err.Error())
	}
}
