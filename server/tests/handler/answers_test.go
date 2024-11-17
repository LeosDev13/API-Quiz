package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quiz-app/server/handler"
	"quiz-app/server/logger"
	"quiz-app/server/repository"
	"quiz-app/shared/dto"
	"testing"
)

func TestSubmitAnswers(t *testing.T) {
	questionRepo := repository.NewInMemoryQuestionRepository()
	leaderboardRepo := repository.NewInMemoryLeaderboardRepository()

	log := logger.New()
	handler := handler.NewSubmitAnswersHandler(questionRepo, leaderboardRepo, log)

	submitRequest := dto.SubmitAnswersRequest{
		Username: "testuser",
		Answers: map[string]string{
			"a1f5a1a2-1234-4d4a-bbbb-55c1a4d3e5f6": "HyperText Transfer Protocol",
			"b2f3b2a3-2345-4e5b-cccc-66d2b5e4f6g7": "Integrated Development Environment",
		},
	}

	requestBody, err := json.Marshal(submitRequest)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	req := httptest.NewRequest("POST", "/answers", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.SubmitAnswers(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %v", status)
	}

	var response dto.SubmitAnswersResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.CorrectAnswers != 2 {
		t.Errorf("Expected 2 correct answers, got %v", response.CorrectAnswers)
	}

	leaderboard := leaderboardRepo.GetLeaderboard()
	if len(leaderboard) != 1 {
		t.Fatalf("Expected 1 entry in the leaderboard, got %d", len(leaderboard))
	}

	if leaderboard[0].Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", leaderboard[0].Username)
	}

	if leaderboard[0].Score != 2 {
		t.Errorf("Expected score 2 for user 'testuser', got %d", leaderboard[0].Score)
	}
}
