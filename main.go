package main

import (
	"net/http"
	"quiz-app/logger"
	"quiz-app/middleware"
	"quiz-app/repository"
	"quiz-app/router"
)

func main() {
	log := logger.New()

	questionRepo := repository.NewMemoryQuestionRepository()
	r := router.NewRouter(questionRepo)

	middlewares := middleware.ApplyMiddlewares(r, log)
	log.Info("Starting server on :3000", nil)
	if err := http.ListenAndServe(":3000", middlewares); err != nil {
		log.Error(err, "Server failed")
	}
}
