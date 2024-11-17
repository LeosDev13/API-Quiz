package repository

import "quiz-app/server/model"

type memoryAnswerRepository struct {
	answers []model.Answer
}

func NewMemoryAnswerRepository() AnswerRepository {
	return &memoryAnswerRepository{}
}

func (r *memoryAnswerRepository) SaveAnswers(answers []model.Answer) error {
	r.answers = append(r.answers, answers...)
	return nil
}

func (r *memoryAnswerRepository) GetAllAnswers() []model.Answer {
	return r.answers
}
