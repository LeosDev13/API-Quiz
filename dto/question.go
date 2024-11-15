package dto

type Question struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}
