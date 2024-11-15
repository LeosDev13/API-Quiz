package router

import (
	"net/http"
	"quiz-app/api"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", api.GetQuestions)
	return mux
}
