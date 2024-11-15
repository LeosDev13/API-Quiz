package api

import (
	"encoding/json"
	"net/http"
)

type Question struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := []Question{
		{
			ID:       "a1f5a1a2-1234-4d4a-bbbb-55c1a4d3e5f6",
			Question: "What does 'HTTP' stand for?",
			Options:  []string{"HyperText Transfer Protocol", "HighText Transfer Protocol", "Hyper Transfer Text Protocol", "HyperTransfer Text Protocol"},
		},
		{
			ID:       "b2f3b2a3-2345-4e5b-cccc-66d2b5e4f6g7",
			Question: "In programming, what does 'IDE' stand for?",
			Options:  []string{"Integrated Development Environment", "Interactive Development Environment", "Integrated Design Environment", "Intelligent Development Environment"},
		},
		{
			ID:       "c3e1c3b4-3456-5f6c-dddd-77e3c6f5g7h8",
			Question: "Which programming language is known as the 'mother of all languages'?",
			Options:  []string{"C", "Assembly", "Python", "Java"},
		},
		{
			ID:       "d4g2d4c5-4567-6g7d-eeee-88f4d7g6h8i9",
			Question: "Which HTML element is used for the largest heading?",
			Options:  []string{"<h1>", "<h6>", "<header>", "<heading>"},
		},
		{
			ID:       "e5h3e5d6-5678-7h8e-ffff-99g5h8i7j9k0",
			Question: "In version control, what does 'Git' help with?",
			Options:  []string{"File Versioning", "File Compression", "File Encryption", "File Editing"},
		},
		{
			ID:       "f6i4f6e7-6789-8i9f-aaaa-00h6i9j8k0l1",
			Question: "What symbol is used to denote comments in Python?",
			Options:  []string{"#", "//", "/*", "!"},
		},
		{
			ID:       "g7j5g7f8-7890-9j0g-bbbb-11i7j0k1l2m3",
			Question: "Which of these is not a programming language?",
			Options:  []string{"HTML", "Python", "Ruby", "Java"},
		},
		{
			ID:       "h8k6h8g9-8901-0k1h-cccc-22j8k1l2m3n4",
			Question: "What does 'CSS' stand for in web development?",
			Options:  []string{"Cascading Style Sheets", "Creative Style System", "Computer Styling Syntax", "Cascading System Styles"},
		},
		{
			ID:       "i9l7i9h0-9012-1l2i-dddd-33k9l2m3n4o5",
			Question: "What is the primary purpose of the 'RAM' in a computer?",
			Options:  []string{"Temporary data storage", "Permanent data storage", "Graphics rendering", "Data encryption"},
		},
		{
			ID:       "j0m8j0i1-0123-2m3j-eeee-44l0m3n4o5p6",
			Question: "Which data structure uses 'LIFO' order?",
			Options:  []string{"Stack", "Queue", "Array", "Tree"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}
