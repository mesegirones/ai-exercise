package domain

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	// Options  []QuestionOptions `json:"options"`
}
