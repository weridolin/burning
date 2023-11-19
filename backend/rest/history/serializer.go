package history

import "github.com/weridolin/burning/models"

type TrainHistoryContentItem struct {
	TrainHistory models.TrainingHistory         `json:"train_history"`
	TrainContent []models.TrainingContentDetail `json:"train_content"`
}
