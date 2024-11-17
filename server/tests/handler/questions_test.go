package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quiz-app/server/handler"
	"quiz-app/server/model"
	"quiz-app/server/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllQuestions_E2E(t *testing.T) {
	repo := repository.NewInMemoryQuestionRepository()

	handler := handler.NewQuestionHandler(repo)

	server := httptest.NewServer(http.HandlerFunc(handler.GetAllQuestions))
	defer server.Close()

	resp, err := http.Get(server.URL + "/questions")
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var questions []model.Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(questions))
}
