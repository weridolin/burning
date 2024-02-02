package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LocalDate time.Time

func (t LocalDate) Format(format string) string {
	//do your serializing here
	stamp := time.Time(t).Format(format)
	// fmt.Println("marshal json -> ", stamp)
	return stamp
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := time.Time(t).Format("2006-01-02")
	// fmt.Println("marshal json -> ", stamp)
	return []byte(fmt.Sprintf(`"%s"`, stamp)), nil
}

func (t *LocalDate) UnmarshalJSON(data []byte) error {
	//do your deserializing here
	// fmt.Println("unmarshal json -> ", string(data))
	// now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	now, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)

	*t = LocalDate(now)
	return err
}

func (t LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	fmt.Println("value -> ", tlt)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalDate) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)

}

type LocalDateTime time.Time

func (t LocalDateTime) Format(format string) string {
	//do your serializing here
	stamp := time.Time(t).Format(format)
	// fmt.Println("marshal json -> ", stamp)
	return stamp
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := time.Time(t).Format("2006-01-02 15:04:05")
	// fmt.Println("marshal json -> ", stamp)
	return []byte(fmt.Sprintf(`"%s"`, stamp)), nil
}

func (t *LocalDateTime) UnmarshalJSON(data []byte) error {
	//do your deserializing here
	// fmt.Println("unmarshal json -> ", string(data))
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	// now, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*t = LocalDateTime(now)
	return err
}

func (t LocalDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	fmt.Println("value -> ", tlt)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalDateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalDateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)

}

type BaseModel struct {
	ID        int            `gorm:"primarykey" json:"id" yaml:"id"`
	CreatedAt *LocalDate     `json:"created_at" yaml:"created_at"`
	UpdatedAt *LocalDate     `json:"updated_at" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" yaml:"deleted_at"`
}

type Setting struct {
	BaseModel
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}
