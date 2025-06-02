package models

type ClazzData struct {
	Id               int    `gorm:"PrimaryKey"`
	DataClazzName    string `gorm:"column:data_clazz_name"`
	DataClazzRoom    string `gorm:"column:data_clazz_room"`
	DataClazzTeacher string `gorm:"column:data_clazz_teacher"`
	DataClazzStatus  string `gorm:"column:data_clazz_status"`
}
