package models

import "gorm.io/gorm"

//个人资料
type PersonProfile struct {
	BaseModel
	UserID                int     `json:"user_id" yaml:"user_id" comment:"用户ID" gorm:"uniqueIndex:idx_user_id"`
	Height                int     `json:"height" yaml:"height" comment:"身高"`
	BodyFatRate           float32 `json:"body_fat_rate" yaml:"body_fat_rate" comment:"体脂率"`
	ChestCircumference    int     `json:"chest_circumference" yaml:"chest_circumference" comment:"胸围"`
	UpperArmCircumference int     `json:"upper_arm_circumference" yaml:"upper_arm_circumference" comment:"上臂围"`
	WaistLine             int     `json:"waistline" yaml:"waistline" comment:"腰围" gorm:"column:waistline"`
	ThighCircumference    int     `json:"thigh_circumference" yaml:"thigh_circumference" comment:"大腿围"`
	CalfCircumference     int     `json:"calf_circumference" yaml:"calf_circumference" comment:"小腿围"`
	HipLine               int     `json:"hipline" yaml:"hipline" comment:"臀围" gorm:"column:hipline"`
	Weight                int     `json:"weight" yaml:"weight" comment:"体重"`
	ShoulderBreadth       int     `json:"shoulder_breadth" yaml:"shoulder_breadth" comment:"肩宽"`
	Uuid                  string  `json:"uuid" yaml:"uuid" comment:"uuid" gorm:"column:uuid"`
	Days                  int     `json:"days" yaml:"days" comment:"签到天数"`
}

func GetUserProfile(user_id int, db *gorm.DB) (profile PersonProfile, err error) {
	var _profile PersonProfile
	err = db.Where("user_id = ?", user_id).First(&_profile).Error
	return _profile, err
}

func UpdateUserProfile(user_id int, params map[string]interface{}, db *gorm.DB) error {
	return db.Model(&PersonProfile{}).Where("user_id = ?", user_id).Updates(params).Error
}

func CreateUserProfile(profile *PersonProfile, db *gorm.DB) error {
	return db.Create(profile).Error
}
