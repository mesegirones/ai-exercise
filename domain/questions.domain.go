package domain

type LLMInput struct {
	Prompt    string `json:"prompt"`
	UserInput string `json:"user_input"`
}

type QuestionInput struct {
	UserID   string `json:"user_id"`
	Question string `json:"question"`
}

type QuestionResponse struct {
	Channel chan QuestionStatus `json:"chanel"`
}

type QuestionStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
