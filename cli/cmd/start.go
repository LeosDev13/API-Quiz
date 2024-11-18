package cmd

import (
	"fmt"
	"log"
	"quiz-app/shared/client"
	"quiz-app/shared/dto"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startQuiz() {
	prompt := promptui.Prompt{
		Label: "Enter your username",
	}

	username, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	apiClient := client.NewAPIClient("http://localhost:3000")

	questions, err := apiClient.FetchQuestions()
	if err != nil {
		log.Fatalf("Error fetching questions: %v", err)
	}

	answers := make([]dto.Answer, 0)

	for _, question := range questions {
		fmt.Printf("\nQuestion: %s\n", question.Question)

		prompt := promptui.Select{
			Label: "Select your answer",
			Items: question.Options,
		}

		_, answer, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
		}

		answers = append(answers, dto.Answer{
			ID:     question.ID,
			Answer: answer,
		})
	}

	result, err := apiClient.SubmitAnswers(answers, username)
	if err != nil {
		log.Fatalf("Error submitting answers: %v", err)
	}

	fmt.Printf("\nYou got %d out of %d questions correct!\n", result.CorrectAnswers, len(questions))
	fmt.Printf("You were better than %.2f%% of all quizzers\n", result.Percentile)
}
