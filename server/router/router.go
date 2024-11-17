package router

import (
	"net/http"
	"quiz-app/server/handler"
	"quiz-app/server/logger"
	"quiz-app/server/repository"
)

func NewRouter(questionRepository repository.QuestionRepository, leaderBoardRepository repository.LeaderboardRepository, log logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", handler.NewQuestionHandler(questionRepository).GetAllQuestions)
	mux.HandleFunc("/answers", handler.NewSubmitAnswersHandler(questionRepository, leaderBoardRepository, log).SubmitAnswers)
	mux.HandleFunc("/leaderboard", handler.NewLeaderboardHandler(leaderBoardRepository, log).GetLeaderboard)
	return mux
}
