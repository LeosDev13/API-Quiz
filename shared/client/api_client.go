package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"quiz-app/shared/dto"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

func (c *APIClient) FetchQuestions() ([]dto.Question, error) {
	resp, err := http.Get(fmt.Sprintf("%s/questions", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var questions []dto.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func (c *APIClient) SubmitAnswers(answers []dto.Answer, username string) (*dto.SubmitAnswersResponse, error) {
	answersMap := make(map[string]string)
	for _, answer := range answers {
		answersMap[answer.ID] = answer.Answer
	}

	req := dto.SubmitAnswersRequest{
		Answers: answersMap, Username: username,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/answers", c.BaseURL), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.SubmitAnswersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) FetchLeaderboard() (*dto.LeaderboardResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/leaderboard", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leaderboard dto.LeaderboardResponse
	if err := json.NewDecoder(resp.Body).Decode(&leaderboard); err != nil {
		return nil, err
	}

	return &leaderboard, nil
}
