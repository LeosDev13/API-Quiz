package main

import (
	"log"
	"net/http"
	"quiz-app/middleware"
	"quiz-app/router"
)

func main() {
	r := router.NewRouter()

	middlewares := middleware.ApplyMiddlewares(r)

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", middlewares); err != nil {
		log.Fatal("Server failed:", err)
	}
}
