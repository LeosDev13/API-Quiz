package router

import (
	"net/http"
	"quiz-app/handler"
	"quiz-app/repository"
)

func NewRouter(repo repository.QuestionRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", handler.NewQuestionHandler(repo).GetAllQuestions)
	mux.HandleFunc("/answers", handler.NewSubmitAnswersHandler(repo).SubmitAnswers)
	return mux
}
