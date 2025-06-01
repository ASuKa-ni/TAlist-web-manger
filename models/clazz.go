package models

import (
	"time"
)

type ClazzData struct {
	Id                  int       `gorm:"PrimaryKey"`
	DataClazzName       string    `gorm:"column:data_clazz_name"`
	DataClazzRoom       string    `gorm:"column:data_clazz_room"`
	DataClazzTeacher    string    `gorm:"column:data_clazz_teacher"`
	DataClazzStartDate  time.Time `gorm:"column:data_clazz_start_date"`
	DataClazzEndDate    time.Time `gorm:"column:data_clazz_end_date"`
	DataClazzStatus     string    `gorm:"column:data_clazz_status"`
	DataClazzLatestDate time.Time `gorm:"column:data_clazz_latest_date"`
}
