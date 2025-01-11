package storage

// import (
// 	"encoding/json"
// 	"ia-exercise/domain"
// 	"os"
// )

// type LoggerProxy interface {
// 	Error(...interface{})
// }

// type StorageConfig struct {
// 	Logger    LoggerProxy
// 	Questions []domain.Question
// 	Answer    []domain.UserAnswer
// }

// func NewStorageCongif(logger LoggerProxy) *StorageConfig {
// 	questions := readFile[[]domain.Question](logger, "assets/questions.json")
// 	answers := readFile[[]domain.UserAnswer](logger, "assets/answers.json")

// 	return &StorageConfig{
// 		Logger:    logger,
// 		Questions: questions,
// 		Answer:    answers,
// 	}
// }

// func readFile[T []domain.Question | []domain.UserAnswer](logger LoggerProxy, fileUrl string) T {
// 	fileData, err := os.ReadFile(fileUrl)
// 	if err != nil {
// 		logger.Error(err)
// 		return nil
// 	}
// 	var parsedData T
// 	err = json.Unmarshal(fileData, &parsedData)
// 	if err != nil {
// 		logger.Error(err)
// 		return nil
// 	}
// 	return parsedData
// }
