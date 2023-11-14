package models

type Actions struct {
	BaseModel
	UUID              string `json:"uuid" yaml:"uuid" comment:"动作UUID" gorm:"uniqueIndex" column:"uuid"`
	ActionsName       string `json:"actions_name" yaml:"actions_name" comment:"动作名称"`
	ActionsType       string `json:"actions_type" yaml:"actions_type" comment:"动作类型"`
	ActionsDesc       string `json:"actions_desc" yaml:"actions_desc" comment:"动作描述"`
	ActionsImg        string `json:"actions_img" yaml:"actions_img" comment:"动作图片"`
	ActionsVideo      string `json:"action_video" yaml:"action_video" comment:"动作视频"`
	ActionsInstrument string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
}
