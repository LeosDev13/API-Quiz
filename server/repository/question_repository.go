package repository

import "quiz-app/server/model"

type QuestionRepository interface {
	GetQuestions() []model.Question
	GetByID(id string) (*model.Question, error)
}
