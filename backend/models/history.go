package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	ActionsName       string `json:"actions_name" yaml:"actions_name" comment:"动作名称"`
	LeftWeight        string `json:"left_weight" yaml:"left_weight" comment:"左侧重量"`
	RightWeight       string `json:"right_weight" yaml:"right_weight" comment:"右侧重量"`
	TotalWeight       string `json:"total_weight" yaml:"total_weight" comment:"总重量"`
	GroupsNumber      string `json:"groups_number" yaml:"groups_number" comment:"组数"`
	UserID            int    `json:"user_id" yaml:"user_id" comment:"用户ID" gorm:"index" column:"user_id"`
	TrainingHistoryId int    `json:"training_history_id" yaml:"training_history_id" comment:"训练历史ID" gorm:"index" column:"training_history_id"`
	ActionsType       string `json:"actions_type" yaml:"actions_type" comment:"动作类型"`
	ActionsInstrument string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
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
}

// 训练模板
type TrainTemplate struct {
	BaseModel
	UserID  int           `json:"user_id" yaml:"user_id" comment:"用户ID"`
	Content TrainsContent `json:"content" yaml:"content" comment:"内容" gorm:"type:jsonb"`
}
