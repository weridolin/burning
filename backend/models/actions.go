package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Actions struct {
	BaseModel
	UUID              string `json:"uuid" yaml:"uuid" comment:"动作UUID" gorm:"uniqueIndex" column:"uuid"`
	ActionsName       string `json:"actions_name" yaml:"actions_name" comment:"动作名称"`
	ActionsType       string `json:"actions_type" yaml:"actions_type" comment:"动作类型"`
	ActionsDesc       string `json:"actions_desc" yaml:"actions_desc" comment:"动作描述"`
	ActionsImg        string `json:"actions_img" yaml:"actions_img" comment:"动作图片"`
	ActionsVideoUri   string `json:"action_video" yaml:"action_video" comment:"动作视频"`
	ActionsInstrument string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
}

func GenerateUUID() string {
	return uuid.New().String()
}

func CreateActions(actions Actions, DB *gorm.DB) (Actions, error) {
	actions.UUID = GenerateUUID()
	err := DB.Create(&actions).Error
	return actions, err
}

func UpdateAction(ID int, params map[string]interface{}, DB *gorm.DB) error {
	var action Actions
	err := DB.Model(&action).Where("id = ?", ID).Updates(params).Error
	return err
}

func QueryActionById(Id int, DB *gorm.DB) (Actions, error) {
	var action Actions
	err := DB.Where("id = ?", Id).First(&action).Error
	return action, err
}

// func (action Actions) Serialize(tx *gorm.DB) (err error) {

// }
func QueryActions(condition map[string]interface{}, DB *gorm.DB) ([]Actions, error) {
	var actions []Actions
	err := DB.Where(condition).Find(&actions).Error
	return actions, err
}
