package actions

type AddActionRequest struct {
	ActionName       string `json:"action_name" yaml:"action_name" comment:"动作名称"`
	ActionType       string `json:"action_type" yaml:"action_type" comment:"动作类型"`
	ActionDesc       string `json:"action_desc" yaml:"action_desc" comment:"动作描述"`
	ActionImg        string `json:"action_img" yaml:"action_img" comment:"动作图片"`
	ActionVideoUri   string `json:"action_video" yaml:"action_video" comment:"动作视频"`
	ActionInstrument string `json:"action_instrument" yaml:"action_instrument" comment:"动作器械"`
}

type UpdateActionRequest struct {
	ID int `json:"id" yaml:"id" uri:"id" comment:"动作ID"`
	AddActionRequest
}
