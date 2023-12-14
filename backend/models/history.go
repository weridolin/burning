package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FoodHistoryContentItem struct {
	Carbon  int `json:"carbon" yaml:"carbon" comment:"碳水化合物(g)"`
	Fat     int `json:"fat" yaml:"fat" comment:"脂肪(g)"`
	Protein int `json:"protein" yaml:"protein" comment:"蛋白质(g)"`
}

type FoodHistoryContent []FoodHistoryContentItem

func (j *FoodHistoryContent) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var result FoodHistoryContent
	err := json.Unmarshal(bytes, &result)
	*j = FoodHistoryContent(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j FoodHistoryContent) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//饮食记录
type FoodHistory struct {
	BaseModel
	UserID  int                `json:"user_id" yaml:"user_id" comment:"用户ID"`
	Content FoodHistoryContent `json:"content" yaml:"content" comment:"内容" gorm:"type:jsonb"`
}

// 训练计划日志

type TrainingContentDetail struct {
	BaseModel
	ActionName        string `json:"action_name" yaml:"action_name" comment:"动作名称"`
	LeftWeight        string `json:"left_weight" yaml:"left_weight" comment:"左侧重量(kg)"`
	RightWeight       string `json:"right_weight" yaml:"right_weight" comment:"右侧重量(kg)"`
	TotalWeight       string `json:"total_weight" yaml:"total_weight" comment:"总重量(kg)"`
	Number            string `json:"number" yaml:"number" comment:"动作次数"`
	UserID            int    `json:"user_id" yaml:"user_id" comment:"用户ID" gorm:"index" column:"user_id"`
	TrainingHistoryId int    `json:"training_history_id" yaml:"training_history_id" comment:"训练历史ID" gorm:"index" column:"training_history_id"`
	ActionType        string `json:"action_type" yaml:"action_type" comment:"动作类型"`
	ActionInstrument  string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
	Finish            bool   `json:"finish" yaml:"finish" comment:"是否完成"`
	ConsumeTime       int    `json:"consume_time" yaml:"consume_time" comment:"消耗时间(s)"`
}

type TrainsContent []TrainingContentDetail

func (j *TrainsContent) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var result TrainsContent
	err := json.Unmarshal(bytes, &result)
	*j = TrainsContent(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j TrainsContent) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//训练历史记录
type TrainingHistory struct {
	BaseModel
	UserID int `json:"user_id" yaml:"user_id" comment:"用户ID" gorm:"index" column:"user_id"`
	// Content TrainsContent `json:"content" yaml:"content" comment:"内容" gorm:"type:jsonb"`
	// ContentDetail string        `json:"content_uuid" yaml:"content_uuid" comment:"内容UUID"`
	Comment   string `json:"comment" yaml:"comment" comment:"评论"`
	TotalTime int    `json:"total_time" yaml:"total_time" comment:"总时间(s)"`
	Title     string `json:"title" yaml:"title" comment:"标题"`
	Finish    bool   `json:"finish" yaml:"finish" comment:"是否完成"`
}

// 训练模板
type TrainTemplate struct {
	BaseModel
	UserID  int           `json:"user_id" yaml:"user_id" comment:"用户ID"`
	Content TrainsContent `json:"content" yaml:"content" comment:"内容" gorm:"type:jsonb"`
}

func CreateTrainingHistory(history TrainingHistory, DB *gorm.DB) (TrainingHistory, error) {
	err := DB.Debug().Create(&history).Error
	return history, err
}

func QueryTrainingHistory(params interface{}, DB *gorm.DB) ([]TrainingHistory, error) {
	var history []TrainingHistory
	err := DB.Debug().Where(params).Find(&history).Error
	return history, err
}

func UpdateTrainHistory(id int, params map[string]interface{}, DB *gorm.DB) error {
	var history TrainingHistory
	err := DB.Model(&history).Where("id = ?", id).Updates(params).Error
	return err
}

func CreateTrainTemplate(template TrainTemplate, DB *gorm.DB) (TrainTemplate, error) {
	err := DB.Create(&template).Error
	return template, err
}

func CreateTrainingContentDetail(detail *TrainingContentDetail, DB *gorm.DB) (*TrainingContentDetail, error) {
	// err := DB.Create(&detail).Error
	// return detail, err
	history, err := QueryTrainingHistory(map[string]interface{}{"id": detail.TrainingHistoryId}, DB)
	if err != nil || len(history) == 0 {
		return nil, err
	}
	detail.TrainingHistoryId = history[0].ID
	err = DB.Create(&detail).Error
	return detail, err
}

func QueryTrainingContentDetail(trainingHistoryId int, DB *gorm.DB) ([]TrainingContentDetail, error) {
	var detail []TrainingContentDetail
	err := DB.Where("training_history_id = ?", trainingHistoryId).Find(&detail).Error
	return detail, err

}

func DeleteTrainingHistory(id int, DB *gorm.DB) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := DB.Unscoped().Where("id = ?", id).Delete(&TrainingHistory{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := DeleteHistoryDetail(map[string]interface{}{"training_history_id": id}, DB); err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

}

func DeleteHistoryDetail(condition map[string]interface{}, DB *gorm.DB) error {
	var detail TrainingContentDetail
	err := DB.Unscoped().Where(condition).Delete(&detail).Error
	return err
}

func UpdateTrainContent(id int, params map[string]interface{}, DB *gorm.DB) error {
	var detail TrainingContentDetail
	err := DB.Model(&detail).Where("id = ?", id).Updates(params).Error
	return err
}

func DeleteTrainingContent(id int, DB *gorm.DB) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := DB.Unscoped().Where("id = ?", id).Delete(&TrainingContentDetail{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func FinishTraining(trainHistory TrainingHistory, contentList []TrainingContentDetail, DB *gorm.DB, user_id int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 先更新训练历史
		fmt.Println("train history", trainHistory)
		err := UpdateTrainHistory(trainHistory.ID, map[string]interface{}{
			"finish":     true,
			"total_time": trainHistory.TotalTime,
			"comment":    trainHistory.Comment,
			"title":      trainHistory.Title,
			"user_id":    user_id,
		}, DB)
		if err != nil {
			return err
		}

		// 批量插入训练内容,range是值拷贝
		if len(contentList) == 0 {
			return nil
		}
		for index, _ := range contentList {
			contentList[index].UserID = user_id
			contentList[index].TrainingHistoryId = trainHistory.ID
		}
		// fmt.Println("content list -> ", contentList, user_id, trainHistory.ID)
		err = DB.Clauses(clause.OnConflict{ //ID 冲突时直接更新
			UpdateAll: true,
			Columns:   []clause.Column{{Name: "id"}},
			// DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
		}).Create(&contentList).Error
		if err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func UpdateTraining(trainHistory TrainingHistory, contentList []TrainingContentDetail, DB *gorm.DB, user_id int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 先更新训练历史
		fmt.Println("train history", trainHistory)
		err := UpdateTrainHistory(trainHistory.ID, map[string]interface{}{
			"finish":     true,
			"total_time": trainHistory.TotalTime,
			"comment":    trainHistory.Comment,
			"title":      trainHistory.Title,
			"user_id":    user_id,
		}, DB)
		if err != nil {
			return err
		}

		// 批量更新内容,range是值拷贝

		// 返回 nil 提交事务
		return nil
	})
}
