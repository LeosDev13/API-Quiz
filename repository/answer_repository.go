package repository

import "quiz-app/model"

type AnswerRepository interface {
	SaveAnswers(answers []model.Answer) error
	GetAllAnswers() []model.Answer
}
