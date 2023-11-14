package models

//个人资料
type PersonProfile struct {
	BaseModel
	UserID                 int     `json:"user_id" yaml:"user_id" comment:"用户ID"`
	Height                 int     `json:"height" yaml:"height" comment:"身高"`
	BodyFatRate            float32 `json:"body_fat_rate" yaml:"body_fat_rate" comment:"体脂率"`
	ChestCircumference     int     `json:"chest_circumference" yaml:"chest_circumference" comment:"胸围"`
	UpperArmCircumference  int     `json:"upper_arm_circumference" yaml:"upper_arm_circumference" comment:"上臂围"`
	AbdominalCircumference int     `json:"abdominal_circumference" yaml:"abdominal_circumference" comment:"腹围"`
	ThighCircumference     int     `json:"thigh_circumference" yaml:"thigh_circumference" comment:"大腿围"`
	CalfCircumference      int     `json:"calf_circumference" yaml:"calf_circumference" comment:"小腿围"`
	HipLine                int     `json:"hip_line" yaml:"hip_line" comment:"臀围"`
	Weight                 int     `json:"weight" yaml:"weight" comment:"体重"`
	ShoulderBreadth        int     `json:"shoulder_breadth" yaml:"shoulder_breadth" comment:"肩宽"`
}
