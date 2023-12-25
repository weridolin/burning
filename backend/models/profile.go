package models

import (
	"gorm.io/gorm"
)

//个人资料
type PersonProfile struct {
	BaseModel
	UserID int    `json:"user_id" yaml:"user_id" comment:"用户ID" gorm:"uniqueIndex:idx_user_id"`
	Uuid   string `json:"uuid" yaml:"uuid" comment:"uuid" gorm:"column:uuid"`
	Days   int    `json:"days" yaml:"days" comment:"累计签到天数"`
}

// 用户签到表
type UserSign struct {
	BaseModel
	UserID int `json:"user_id" yaml:"user_id" comment:"用户ID"`
	Type   int `json:"type" yaml:"type" comment:"签到类型"`
	Reward int `json:"reward" yaml:"reward" comment:"奖励"`
}

// 用户身体数据
type BodyInfo struct {
	BaseModel
	Height                string `json:"height" yaml:"height" gorm:"size:255;comment:身高"`
	BodyFatRate           string `json:"body_fat_rate" yaml:"body_fat_rate" gorm:"size:255;comment:体脂率"`
	ChestCircumference    string `json:"chest_circumference" yaml:"chest_circumference" gorm:"size:255;comment:胸围"`
	UpperArmCircumference string `json:"upper_arm_circumference" yaml:"upper_arm_circumference" gorm:"size:255;comment:上臂围"`
	WaistLine             string `json:"waistline" yaml:"waistline" gorm:"column:waistline;size:255;comment:腰围"`
	ThighCircumference    string `json:"thigh_circumference" yaml:"thigh_circumference" gorm:"size:255;comment:大腿围"`
	CalfCircumference     string `json:"calf_circumference" yaml:"calf_circumference" gorm:"size:255;comment:小腿围"`
	HipLine               string `json:"hipline" yaml:"hipline"  gorm:"column:hipline;size:255;comment:臀围"`
	Weight                string `json:"weight" yaml:"weight" gorm:"size:255;comment:体重"`
	ShoulderBreadth       string `json:"shoulder_breadth" yaml:"shoulder_breadth" gorm:"size:255;comment:肩宽"`
	UserID                int    `json:"user_id" yaml:"user_id" gorm:"comment:用户ID"`
	Date                  string `json:"date" yaml:"date" gorm:"size:255;comment:更新日期"`
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

func Sign(user_id int, db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 签到
		var sign UserSign
		sign.UserID = user_id
		sign.Type = 1
		sign.Reward = 1

		if err := tx.Create(&sign).Error; err != nil {
			return nil
		}

		// 更新累计签到天数
		if err := tx.Model(&PersonProfile{}).Where("user_id = ?", user_id).Update("days", gorm.Expr("days + ?", 1)).Error; err != nil {
			return nil
		}
		return nil
	})
}

func GetLastSignDate(userID int, db *gorm.DB) (date string, err error) {
	var sign UserSign
	err = db.Where("user_id = ?", userID).Order("created_at desc").First(&sign).Error
	if err != nil {
		return "", err
	}
	return sign.CreatedAt.Format("2006-01-02"), nil
}

func GetUserBodyInfo(userID int, Date string, DB *gorm.DB) (BodyInfo, error) {
	var bodyInfo BodyInfo
	err := DB.Where("user_id = ? and date = ?", userID, Date).First(&bodyInfo).Error
	return bodyInfo, err
}

func GetAllUserBodyInfo(userID int, DB *gorm.DB) []BodyInfo {
	var bodyInfos []BodyInfo
	DB.Where("user_id = ?", userID).Find(&bodyInfos)
	return bodyInfos
}

func UpdateBodyInfo(UserID int, Params map[string]interface{}, db *gorm.DB) error {
	return db.Model(&BodyInfo{}).Where("user_id = ? and date = ?", UserID, Params["date"]).Updates(Params).Error
}

func CreateNewBodyInfo(bodyInfo *BodyInfo, db *gorm.DB) error {
	return db.Create(bodyInfo).Error
}
