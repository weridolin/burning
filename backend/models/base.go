package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LocalTime time.Time

func (t LocalTime) Format(format string) string {
	//do your serializing here
	stamp := time.Time(t).Format(format)
	// fmt.Println("marshal json -> ", stamp)
	return stamp
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := time.Time(t).Format("2006-01-02")
	// fmt.Println("marshal json -> ", stamp)
	return []byte(fmt.Sprintf(`"%s"`, stamp)), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	//do your deserializing here
	// fmt.Println("unmarshal json -> ", string(data))
	// now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	now, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)

	*t = LocalTime(now)
	return err
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	fmt.Println("value -> ", tlt)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)

}

// func (t *LocalTime) Format() {

// }

type BaseModel struct {
	ID        int            `gorm:"primarykey" json:"id" yaml:"id"`
	CreatedAt *LocalTime     `json:"created_at" yaml:"created_at"`
	UpdatedAt *LocalTime     `json:"updated_at" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" yaml:"deleted_at"`
}
