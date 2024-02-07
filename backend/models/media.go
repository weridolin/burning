package models

import "gorm.io/gorm"

type Music struct {
	BaseModel
	Title   string `json:"title" yaml:"title" gorm:"comment:音乐标题" `
	FileUrl string `json:"fileUrl" yaml:"fileUrl"  gorm:"comment:音乐文件地址" `
	Image   string `json:"image" yaml:"image"   gorm:"comment:音乐图片"`
	Singer  string `json:"singer" yaml:"singer"    gorm:"comment:歌手"`
	Enable  bool   `json:"enable" yaml:"enable" gorm:"comment:是否启用;default:true"`
}

// 分页查询音乐列表
func QueryMusicList(page, pageSize int, condition map[string]interface{}, DB *gorm.DB) ([]Music, int64, error) {
	var musics []Music
	var total int64
	DB.Model(&Music{}).Where(condition).Count(&total)
	err := DB.Where(condition).Offset((page - 1) * pageSize).Limit(pageSize).Find(&musics).Error
	return musics, total, err
}
