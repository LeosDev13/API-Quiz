package repository

import "quiz-app/model"

type QuestionRepository interface {
	GetQuestions() []model.Question
	GetByID(id string) (*model.Question, error)
}
