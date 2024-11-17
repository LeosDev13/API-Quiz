package main

import (
	"net/http"
	"quiz-app/server/logger"
	"quiz-app/server/middleware"
	"quiz-app/server/repository"
	"quiz-app/server/router"
)

func main() {
	log := logger.New()

	questionRepo := repository.NewInMemoryQuestionRepository()
	leaderboardRepo := repository.NewInMemoryLeaderboardRepository()

	r := router.NewRouter(questionRepo, leaderboardRepo, log)

	middlewares := middleware.ApplyMiddlewares(r, log)

	log.Info("Starting server on :3000", nil)
	if err := http.ListenAndServe(":3000", middlewares); err != nil {
		log.Error(err, "Server failed")
	}
}
