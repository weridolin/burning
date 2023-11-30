package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Actions struct {
	BaseModel
	UUID             string `json:"uuid" yaml:"uuid" comment:"动作UUID" gorm:"uniqueIndex" column:"uuid"`
	ActionName       string `json:"action_name" yaml:"action_name" comment:"动作名称"`
	ActionType       string `json:"action_type" yaml:"action_type" comment:"动作类型"`
	ActionDesc       string `json:"action_desc" yaml:"action_desc" comment:"动作描述"`
	ActionImg        string `json:"action_img" yaml:"action_img" comment:"动作图片"`
	ActionVideoUri   string `json:"action_video" yaml:"action_video" comment:"动作视频"`
	ActionInstrument string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
	ActionEnName     string `json:"action_en_name" yaml:"action_en_name" comment:"动作英文名"`
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
