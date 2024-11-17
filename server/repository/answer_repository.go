package repository

import "quiz-app/server/model"

type AnswerRepository interface {
	SaveAnswers(answers []model.Answer) error
	GetAllAnswers() []model.Answer
}
