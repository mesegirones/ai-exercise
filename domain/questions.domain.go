package domain

type DataTypeEnum string

const (
	DataTypeEnumLANGUAGE DataTypeEnum = "LANGUAGE"
	DataTypeEnumSUMMARY  DataTypeEnum = "SUMMARY"
	DataTypeEnumRESULT   DataTypeEnum = "RESULT"
	DataTypeEnumERROR    DataTypeEnum = "ERROR"
)

type LLMInput struct {
	Prompt   string       `json:"prompt"`
	DataType DataTypeEnum `json:"data_type"`
}

type LLMOutput struct {
	Message  string       `json:"message"`
	DataType DataTypeEnum `json:"data_type"`
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
